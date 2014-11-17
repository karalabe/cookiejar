// CookieJar - A contestant's algorithm toolbox
// Copyright 2014 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The toolbox is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).

package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/karalabe/cookiejar.v2/exts/fmtext"
)

type gameStats struct {
	platinum []int
	owner    map[int]int
	fight    map[int]bool
	fleet    map[int][]int
}

//
func simulate(database string, ais string, user string, players int, threads int) (map[string]int, error) {
	// Open the replay database
	db := make(map[int]*gameDetails)
	if file, err := os.Open(database); err != nil {
		return nil, err
	} else {
		defer file.Close()
		if err := gob.NewDecoder(file).Decode(&db); err != nil {
			return nil, err
		}
	}
	// Load all the pre-set AIs
	agents := []string{}
	files, err := ioutil.ReadDir(ais)
	if err == nil {
		for _, f := range files {
			agents = append(agents, filepath.Join(ais, f.Name()))
		}
	}
	// Battle it out
	scores := make([]uint32, len(agents)+1)

	pend := new(sync.WaitGroup)
	pend.Add(len(db))

	limiter := make(chan struct{}, threads)
	for _, game := range db {
		go func() {
			defer pend.Done()

			limiter <- struct{}{}
			if err := matcher(game, agents, user, players, []int{}, scores); err != nil {
				log15.Crit("Failed to run matchmaker: %v.", err)
			}
			<-limiter
		}()
	}
	pend.Wait()

	results := make(map[string]int)
	for i, ai := range agents {
		results[ai] = int(scores[i])
	}
	results[user] = int(scores[len(scores)-1])

	return results, nil
}

func matcher(game *gameDetails, ais []string, user string, players int, opponents []int, scores []uint32) error {
	// If the match is made, simulate and score
	if len(opponents) == players-1 {
		match := make([]string, players)
		for i, ai := range opponents {
			match[i] = ais[ai]
		}
		match[players-1] = user

		if winner, err := battle(game, match); err != nil {
			return err
		} else {
			if winner < players-1 {
				atomic.AddUint32(&scores[opponents[winner]], 1)
			} else {
				atomic.AddUint32(&scores[len(scores)-1], 1)
			}
		}
		return nil
	}
	// Otherwise get a new player into the match
	for i := 0; i < len(ais); i++ {
		opponents = append(opponents, i)
		if err := matcher(game, ais, user, players, opponents, scores); err != nil {
			return err
		}
		opponents = opponents[:len(opponents)-1]
	}
	return nil
}

// Runs a single battle between players on a given board.
func battle(game *gameDetails, players []string) (int, error) {
	log15.Info("Running battle", "ais", players)

	// Create the game stats for the current battle
	stats := &gameStats{
		platinum: []int{200, 200, 200, 200},
		owner:    make(map[int]int),
		fight:    make(map[int]bool),
		fleet:    make(map[int][]int),
	}
	for i := 0; i < game.Zones; i++ {
		stats.owner[i] = -1
		stats.fight[i] = false
		stats.fleet[i] = []int{0, 0, 0, 0}
	}
	// Create the AI processes and attach to their streams
	ins := make([]io.Writer, len(players))
	outs := make([]io.Reader, len(players))
	cmds := make([]*exec.Cmd, len(players))
	for i, player := range players {
		cmds[i] = exec.Command(player)
		ins[i], _ = cmds[i].StdinPipe()
		outs[i], _ = cmds[i].StdoutPipe()
	}
	// Start each of the processes and initialize them
	for i := 0; i < len(players); i++ {
		if err := cmds[i].Start(); err != nil {
			return -1, err
		}
		defer cmds[i].Process.Kill()

		fmt.Fprintf(ins[i], "%d %d %d %d\n", len(players), i, game.Zones, game.Peers)
		for id, plat := range game.Mines {
			fmt.Fprintf(ins[i], "%d %d\n", id, plat)
		}
		for src, peers := range game.Links {
			for _, dst := range peers {
				fmt.Fprintf(ins[i], "%d %d\n", src, dst)
			}
		}
	}
	// Iterate the game until a winning condition is reached
	for r := 0; r < 200; r++ {
		// Distribution: add the mined platinum to each players assets
		for id, plat := range game.Mines {
			if owner := stats.owner[id]; owner != -1 {
				stats.platinum[owner] += plat
			}
		}
		// Moving and buying: update each player and fetch the moves
		moves, deploys := make([]string, len(players)), make([]string, len(players))
		for i := 0; i < len(players); i++ {
			fmt.Fprintf(ins[i], "%d\n", stats.platinum[i])
			for zone := 0; zone < game.Zones; zone++ {
				fmt.Fprintf(ins[i], "%d %d %d %d %d %d\n", zone, stats.owner[zone], stats.fleet[zone][0], stats.fleet[zone][1], stats.fleet[zone][2], stats.fleet[zone][3])
			}
			moves[i], deploys[i] = fmtext.FscanLine(outs[i]), fmtext.FscanLine(outs[i])
		}
		for i := 0; i < len(players); i++ {
			if moves[i] == "WAIT" {
				continue
			}
			in := strings.NewReader(moves[i])
			for {
				// Fetch the move request
				var count, src, dst int
				if _, err := fmt.Fscan(in, &count, &src, &dst); err != nil {
					break
				}
				// Validate the move
				if stats.fleet[src][i] < count {
					log15.Warn("Not enough pods", "zone", src, "available", stats.fleet[src][i], "moved", count)
					continue
				}
				if stats.fight[src] && stats.owner[dst] != i && stats.owner[dst] != -1 {
					log15.Warn("Invalid flee destination", "zone", dst, "owner", stats.owner[dst])
					continue
				}
				// Update the fleet stats
				stats.fleet[src][i] -= count
				stats.fleet[dst][i] += count
			}
		}
		for i := 0; i < len(players); i++ {
			if deploys[i] == "WAIT" {
				continue
			}
			in := strings.NewReader(deploys[i])
			for {
				// Fetch the deploy request
				var count, dst int
				if _, err := fmt.Fscan(in, &count, &dst); err != nil {
					break
				}
				// Validate the deploy
				if count*20 > stats.platinum[i] {
					log15.Warn("Not enough platinum", "available", stats.platinum[i], "spent", count*20)
					continue
				}
				if stats.owner[dst] != i && stats.owner[dst] != -1 {
					log15.Warn("Invalid deploy destination", "zone", dst, "owner", stats.owner[dst])
					continue
				}
				// Update the fleet stats
				stats.platinum[i] -= count * 20
				stats.fleet[dst][i] += count
			}
		}
		// Fighting: kill of a max of three pods on each zone
		for id := 0; id < game.Zones; id++ {
			for fight := 0; fight < 3; fight++ {
				owner, die := -1, false
				for i := 0; i < len(players); i++ {
					if stats.fleet[id][i] != 0 {
						if owner == -1 {
							owner = i
						} else {
							die = true
							break
						}
					}
				}
				if die {
					for i := 0; i < len(players); i++ {
						if stats.fleet[id][i] > 0 {
							stats.fleet[id][i]--
						}
					}
				}
			}
		}
		// Owning: sort out the last standing pods
		for id := 0; id < game.Zones; id++ {
			owner := -1
			for i := 0; i < len(players); i++ {
				if stats.fleet[id][i] != 0 {
					if owner == -1 {
						owner = i
					} else {
						owner = -2
						break
					}
				}
			}
			if owner >= 0 {
				stats.owner[id] = owner
			}
			stats.fight[id] = (owner == -2)
		}
		// Check for winning conditions
	}
	// Report the winner
	zones := make([]int, len(players))
	for id := 0; id < game.Zones; id++ {
		if owner := stats.owner[id]; owner != -1 {
			zones[owner]++
		}
	}
	best := 0
	for i := 1; i < len(players); i++ {
		if zones[best] < zones[i] {
			best = i
		}
	}
	return best, nil
}
