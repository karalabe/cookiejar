package main

import (
	"flag"
	"fmt"
	"runtime"
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

var ais = flag.String("ais", "./ais", "Folder containing pre-selected AIs")
var user = flag.String("user", "./user", "Player AI agent to evaluate")
var players = flag.Int("players", 2, "Number of players to simulate")
var threads = flag.Int("threads", 2, "Number of simulations to run in parallel")

func main() {
	runtime.GOMAXPROCS(16)
	flag.Parse()

	// If additional matches are needed, crawl them
	if *crawl > 0 {
		if err := update(*boards, *crawl, *sleep); err != nil {
			fmt.Printf("Failed to crawl replays: %v.\n", err)
			return
		}
	} else {
		res, err := simulate(*boards, *ais, *user, *players, *threads)
		if err != nil {
			fmt.Printf("Failed to run simulation: %v.\n", err)
			return
		}
		for ai, score := range res {
			fmt.Printf("%40s: %d wins\n", ai, score)
		}
	}
}
