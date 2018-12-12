package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/bakerolls/adventofcode-2018/day-11/grid"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	serial, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		log.Fatal(err)
	}

	// This is unreasonable slow.
	width, height := 300, 300
	var max, maxX, maxY, maxSize int
	for s := 0; s < width; s++ {
		g := grid.Grid(serial)
		x, y, m := g.Max(width, height, s)
		if m > max {
			max, maxX, maxY, maxSize = m, x, y, s
		}
		fmt.Printf("%d,%d,%d,%d,%d\n", s, maxX, maxY, maxSize, max)
	}
	fmt.Printf("%d,%d,%d\n", maxX, maxY, maxSize)
}
