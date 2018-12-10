package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"sort"

	"github.com/BakeRolls/adventofcode-2018/day-06/grid"
)

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

	g := grid.New()
	for i, p := range xys {
		g.SetPoint(p, i+1)
	}

	total := int(math.Max(float64(maxX-minX), float64(maxY-minY)))
	for r := 1; r < total; r++ {
		g.Expand(r)
	}

	// padding is an arbitrary number to make the plot look a little nicer.
	padding := 10

	// Find all points that are directly at an edge of the grid. This is not the
	// a very good solution, but it worls (if expend was called enough times).
	b := image.Rect(minX-padding, minY-padding, maxX+padding, maxY+padding)
	edgePoints := map[int]bool{}
	for x := b.Min.X; x < b.Max.X; x++ {
		edgePoints[g.At(grid.XY{x, b.Min.Y})] = true
		edgePoints[g.At(grid.XY{x, b.Max.Y})] = true
	}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		edgePoints[g.At(grid.XY{b.Min.X, y})] = true
		edgePoints[g.At(grid.XY{b.Max.X, y})] = true
	}

	fmt.Println(g.Max(edgePoints))

	w, err := os.Create("grid.png")
	if err != nil {
		log.Fatal(err)
	}
	img := g.Image(b)
	if err := png.Encode(w, img); err != nil {
		log.Fatal(err)
	}
}
