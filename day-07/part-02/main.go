package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/bakerolls/adventofcode-2018/day-07/node"
)

var (
	numWorkers = 5
	seconds    = 60
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	nodes, err := node.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	var current node.Nodes
	currentMap := map[rune]bool{}
	for _, n := range nodes {
		if len(n.Ancestors) == 0 {
			current = append(current, n)
			currentMap[n.Name] = true
		}
	}

	var output node.Nodes
	done := map[rune]bool{}
	workers := make(workers, numWorkers)
	for s := 0; ; s++ {
		for i := range workers {
			workers[i].steps--
		}

		for i, w := range workers {
			if w.steps > 0 || w.node == nil {
				continue
			}
			done[w.node.Name] = true
			output = append(output, w.node)
			for n := range w.node.Descendants {
				if !done[n] && !currentMap[n] {
					current = append(current, nodes[n])
					currentMap[n] = true
				}
			}
			workers[i].node = nil
		}

		sort.Sort(current)
		for i, w := range workers {
			if len(current) == 0 {
				break
			}
			if w.node != nil {
				continue
			}
			for j, c := range current {
				if !c.Ready(done) {
					continue
				}
				workers[i].steps = int(c.Name) - int('A') + 1 + seconds
				workers[i].node = c
				current = append(current[:j], current[j+1:]...)
				break
			}
		}

		fmt.Printf("%4d ", s)
		for _, w := range workers {
			if w.node == nil {
				fmt.Printf("%2s ", ".")
				continue
			}
			fmt.Printf("%2s ", w.node)
		}
		for _, n := range output {
			fmt.Printf("%s", n)
		}
		fmt.Println()

		if workers.done(done) {
			break
		}
	}
}

type sortRunes []rune

func (rs sortRunes) Len() int           { return len(rs) }
func (rs sortRunes) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs sortRunes) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
