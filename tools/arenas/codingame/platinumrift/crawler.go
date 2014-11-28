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

// Contains the replay crawler to scoop up unseen game boards for simulations.

package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"gopkg.in/inconshreveable/log15.v2"
)

func update(database string, crawl int, sleep int) error {
	// Create a reusable http client
	client := new(http.Client)

	// Open the replay database
	db := make(map[int]*gameDetails)
	file, err := os.Open(database)
	if err == nil {
		if err := gob.NewDecoder(file).Decode(&db); err != nil {
			return err
		}
		file.Close()
	}
	log15.Info("Updating database", "games", len(db), "add", crawl)

	// Retrieve the current leader board
	log15.Info("Crawling leader board")
	players, err := leaderboard(client)
	if err != nil {
		return err
	}
	// For each player, retrieve the replay listing
	for _, playerId := range players {
		// Stop enough replays were retrieved
		if crawl <= 0 {
			break
		}
		log15.Info("Crawling player replay list", "playerId", playerId)
		games, err := replays(client, playerId)
		if err != nil {
			return err
		}
		// For each replay, retrieve the game details
		for _, gameId := range games {
			// Stop enough replays were retrieved
			if crawl <= 0 {
				break
			}
			// Skip any previously retrieved games
			if _, ok := db[gameId]; ok {
				continue
			}
			// Fetch the game details and store it
			log15.Info("Crawling replay details", "gameId", gameId)
			game, err := details(client, gameId)
			if err != nil {
				return err
			}
			db[gameId], crawl = game, crawl-1

			// Sleep a bit not to overload the remote server
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}
	// Update the replay database and return
	file, err = os.OpenFile(database, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	log15.Info("Database updated", "games", len(db))
	return gob.NewEncoder(file).Encode(db)
}

// Retrieves the current leader board.
func leaderboard(client *http.Client) ([]int, error) {
	// Retrieve the leader board listing
	url := "http://www.codingame.com/services/PlayersAgentsRemoteService/findAllValidByChallengePublicId"
	res, err := client.Post(url, "application/json", strings.NewReader("['platinum-rift']"))
	if err != nil {
		return nil, err
	}
	// Define the schema and extract the needed fields
	type entry struct {
		CandidateId int `json:"candidateId"`
	}
	listing := new(struct {
		Players []*entry `json:"success"`
	})
	if err := json.NewDecoder(res.Body).Decode(listing); err != nil {
		return nil, err
	}
	// Flatten the player ids
	players := make([]int, len(listing.Players))
	for i, player := range listing.Players {
		players[i] = player.CandidateId
	}
	return players, nil
}

// Retrieves the list of replays of a given player/
func replays(client *http.Client, player int) ([]int, error) {
	// Retrieve the replay listing
	url := "http://www.codingame.com/services/gamesPlayersRankingRemoteService/findAllByUserId"
	res, err := client.Post(url, "application/json", strings.NewReader(fmt.Sprintf("[%d]", player)))
	if err != nil {
		return nil, err
	}
	// Define the schema and extract the needed fields
	type entry struct {
		GameId int `json:"gameId"`
	}
	listing := new(struct {
		Replays []*entry `json:"success"`
	})
	if err := json.NewDecoder(res.Body).Decode(listing); err != nil {
		return nil, err
	}
	// Flatten the replay ids
	games := make([]int, len(listing.Replays))
	for i, game := range listing.Replays {
		games[i] = game.GameId
	}
	return games, nil
}

// Retrieves the game details of a particular replay.
func details(client *http.Client, id int) (*gameDetails, error) {
	// Retrieve the replay details
	url := "http://www.codingame.com/services/gameResultRemoteService/findInformationByIdAndSaveGame"
	res, err := client.Post(url, "application/json", strings.NewReader(fmt.Sprintf("[%d, -999]", id)))
	if err != nil {
		return nil, err
	}
	// Define the schema and extract the needed fields
	type gameResult struct {
		Input string   `json:"uinput"`
		Views []string `json:"views"`
	}
	type reply struct {
		GameResult gameResult `json:"gameResult"`
	}
	result := new(struct {
		Reply *reply `json:"success"`
	})
	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return nil, err
	}

	game := &gameDetails{
		Links: make(map[int][]int),
		Mines: make(map[int]int),
	}
	// Extract the game board layout
	layout := strings.NewReader(regexp.MustCompile("map=([0-9 ]+)").FindStringSubmatch(result.Reply.GameResult.Input)[1])
	board := make(map[int][]int)
	for {
		var col, row int
		if _, err := fmt.Fscan(layout, &col, &row); err != nil {
			break
		}
		// Inject the hexagon into the game grid
		column := board[col]
		for len(column) < row+1 {
			column = append(column, -1)
		}
		column[row] = game.Zones
		board[col], game.Zones = column, game.Zones+1
	}
	// Initialize the board structure
	for i := 0; i < len(board); i++ {
		for j, id := range board[i] {
			if id == -1 {
				continue
			}
			// Inject a link to all previous column zones
			if i > 0 {
				prev, first := board[i-1], j-(i+1)%2
				for k := 0; k < 2; k++ {
					cell := first + k
					if cell > 0 && cell < len(prev) && prev[cell] != -1 {
						game.Links[id] = append(game.Links[id], prev[cell])
						game.Peers++
					}
				}
			}
		}
	}
	// Extract the platinum distribution
	input := bufio.NewReader(strings.NewReader(result.Reply.GameResult.Views[0]))

	input.ReadString('\n')
	input.ReadString('\n')
	input.ReadString('\n')
	input.ReadString('\n')

	for i := 0; i < game.Zones; i++ {
		var col, row, plat, id int
		line, err := input.ReadString('\n')
		if err != nil {
			return nil, err
		}
		fmt.Sscan(line, &col, &row, &plat, &id)
		game.Mines[id] = plat
	}
	return game, nil
}
