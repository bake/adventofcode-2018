package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bakerolls/adventofcode-2018/day-09/game"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var numPlayers, numMarbles int
	_, err = fmt.Sscanf(string(b), "%d players; last marble is worth %d points", &numPlayers, &numMarbles)
	if err != nil {
		log.Fatal(err)
	}

	// It's fast enough on my machine.
	g := game.New(numPlayers, numMarbles*100)
	for g.Round() {
	}

	var max int
	for _, points := range g.Players {
		if points > max {
			max = points
		}
	}
	fmt.Println(max)
}
