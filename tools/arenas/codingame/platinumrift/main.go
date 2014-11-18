package main

import (
	"flag"
	"fmt"
	"runtime"

	"gopkg.in/inconshreveable/log15.v2"
)

type gameDetails struct {
	Zones int
	Peers int
	Links map[int][]int
	Mines map[int]int
}

var boards = flag.String("boards", "boards.db", "Database file containing real game boards")
var crawl = flag.Int("crawl", 0, "Number of fresh replays to crawl")
var sleep = flag.Int("sleep", 0, "Milliseconds to sleep between crawls")

var ais = flag.String("ai", "./aibin", "Folder containing pre-selected AIs")
var user = flag.String("user", "./user", "Player AI agent to evaluate")
var players = flag.Int("players", 2, "Number of players to simulate")
var threads = flag.Int("threads", 0, "Concurrent simulations (default = #cores)")

func main() {
	log15.Root().SetHandler(log15.LvlFilterHandler(log15.LvlInfo, log15.StderrHandler))
	runtime.GOMAXPROCS(2 * runtime.NumCPU())

	flag.Parse()
	if *threads == 0 {
		*threads = runtime.NumCPU()
	}
	// If additional matches are needed, crawl them
	if *crawl > 0 {
		if err := update(*boards, *crawl, *sleep); err != nil {
			fmt.Printf("Failed to crawl replays: %v.\n", err)
			return
		}
	} else {
		ais, scores, err := simulate(*boards, *ais, *user, *players, *threads)
		if err != nil {
			fmt.Printf("Failed to run simulation: %v.\n", err)
			return
		}
		for i := 0; i < len(ais); i++ {
			fmt.Printf("%40s: %d wins\n", ais[i], scores[i])
		}
	}
}
