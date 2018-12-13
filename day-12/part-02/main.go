package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/bakerolls/adventofcode-2018/day-12/pots"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	state := pots.Parse(lines[0])
	rules := []pots.Rule{}
	for _, raw := range lines[1:] {
		del := strings.Index(raw, "=>")
		if del < 0 {
			continue
		}
		rules = append(rules, pots.Rule{
			Pots:   pots.Parse(raw[:del]),
			Result: pots.Parse(raw[del+2:])[0],
		})
	}

	// Total number of generations. This example only evaluates the first 98
	// generations until it extrapolates.
	numGens := 50000000000

	// Expand the pot slice in both directions (even though I know that this
	// automata only moves to the right). The number by whith the slice expends
	// has to be >= the number of evaluated generations. This could happen inside
	// the loop.
	expand := 500
	state = append(state, make(pots.Pots, expand)...)
	state = append(make(pots.Pots, expand), state...)

	var gen int
	prev := make(pots.Pots, len(state))
	for gen = 1; gen < numGens; gen++ {
		state = state.Evolve(rules)

		// Evaluate the state until it is the previous one shifted one pot to the
		// right.
		// The clarify this, plaese see generations.png. It is a plot of the first
		// 200 generations.
		if prev.Equal(state[1:]) {
			break
		}
		prev = state
	}

	// Calculate the extrapolated sum.
	var sum int
	for i, p := range state {
		if !p {
			continue
		}
		sum += numGens - gen + i - expand
	}
	fmt.Println(sum)
}
