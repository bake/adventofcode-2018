package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/bakerolls/adventofcode-2018/day-07/node"
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

	done := map[rune]bool{}
	curr := map[rune]bool{}
	for name, n := range nodes {
		if n.Ready(done) {
			curr[name] = true
		}
	}

	for len(done) < len(nodes) {
		if len(curr) == 0 {
			break
		}

		var currs []rune
		for name := range curr {
			currs = append(currs, name)
		}
		sort.Sort(sortRunes(currs))

		for _, name := range currs {
			if !nodes[name].Ready(done) {
				continue
			}
			delete(curr, name)
			done[name] = true
			for name := range nodes[name].Descendants {
				curr[name] = true
			}
			fmt.Print(string(name))
			break
		}
	}
	fmt.Println()
}

type sortRunes []rune

func (rs sortRunes) Len() int           { return len(rs) }
func (rs sortRunes) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs sortRunes) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
