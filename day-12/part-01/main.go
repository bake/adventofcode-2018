package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
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

	numGens := 20

	state = append(state, make(pots.Pots, numGens)...)
	state = append(make(pots.Pots, numGens), state...)

	img := image.NewGray(image.Rect(0, 0, len(state), numGens))

	for g := 0; g < numGens; g++ {
		state = state.Evolve(rules)
		for i, p := range state {
			if !p {
				continue
			}
			img.Set(i, g, color.White)
		}
	}

	var sum int
	for i, p := range state {
		if p {
			sum += i - numGens
		}
	}
	fmt.Println(sum)

	w, err := os.Create("generations.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(w, img); err != nil {
		log.Fatal(err)
	}
}
