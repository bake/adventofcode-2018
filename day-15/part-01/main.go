package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/bakerolls/adventofcode-2018/day-15/game"
	"github.com/bakerolls/adventofcode-2018/day-15/pathfinding"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := bytes.Split(bytes.TrimSpace((b)), []byte("\n"))
	g := game.New(grid, pathfinding.BreadthFirstSearch)
	for i := 0; g.Turn(); i++ {
		fmt.Printf("After %d rounds:\n", i+1)
		fmt.Println(g)
		time.Sleep(time.Second / 10)
	}
	turns, hp := g.Outcome()
	fmt.Println(turns, hp, turns*hp)
}
