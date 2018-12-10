package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"sort"

	"github.com/BakeRolls/adventofcode-2018/day-06/grid"
)

var maxDist = 10000

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var xys []grid.XY
	s := bufio.NewScanner(f)
	for s.Scan() {
		var p grid.XY
		_, err := fmt.Sscanf(s.Text(), "%d,%d", &p[0], &p[1])
		if err != nil {
			log.Fatal(err)
		}
		xys = append(xys, p)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(grid.SortByX(xys))
	minX, maxX := xys[0].X(), xys[len(xys)-1].X()
	sort.Sort(grid.SortByY(xys))
	minY, maxY := xys[0].Y(), xys[len(xys)-1].Y()
	b := image.Rect(minX, minY, maxX, maxY)

	// I'm tired, let me just bruteforce this.
	var locs []grid.XY
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if dist(grid.XY{x, y}, xys) >= maxDist {
				continue
			}
			locs = append(locs, grid.XY{x, y})
		}
	}
	fmt.Println(len(locs))
}

func dist(p grid.XY, xys []grid.XY) int {
	var dist int
	for _, q := range xys {
		dist += int(math.Abs(float64(p.X()-q.X()))) + int(math.Abs(float64(p.Y()-q.Y())))
		if dist >= maxDist {
			break
		}
	}
	return dist
}
