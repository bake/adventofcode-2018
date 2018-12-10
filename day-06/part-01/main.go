package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var xys []xy
	s := bufio.NewScanner(f)
	for s.Scan() {
		var p xy
		_, err := fmt.Sscanf(s.Text(), "%d,%d", &p[0], &p[1])
		if err != nil {
			log.Fatal(err)
		}
		xys = append(xys, p)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sortByX(xys))
	minX, maxX := xys[0].x(), xys[len(xys)-1].x()
	sort.Sort(sortByY(xys))
	minY, maxY := xys[0].y(), xys[len(xys)-1].y()

	g := newGrid()
	for i, p := range xys {
		g.setPoint(p, i+1)
	}

	for r := 1; r < 1000; r++ {
		g.expand(r)
	}

	// padding is an arbitrary number to make the plot look a little nicer.
	padding := 10

	// Find all points that are directly at an edge of the grid. This is not the
	// a very good solution, but it worls (if expend was called enough times).
	b := image.Rect(minX-padding, minY-padding, maxX+padding, maxY+padding)
	edgePoints := map[int]bool{}
	for x := b.Min.X; x < b.Max.X; x++ {
		edgePoints[g.at(xy{x, b.Min.Y})] = true
		edgePoints[g.at(xy{x, b.Max.Y})] = true
	}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		edgePoints[g.at(xy{b.Min.X, y})] = true
		edgePoints[g.at(xy{b.Max.X, y})] = true
	}

	fmt.Println(g.max(edgePoints))

	w, err := os.Create("grid.png")
	if err != nil {
		log.Fatal(err)
	}
	img := g.image(b)
	if err := png.Encode(w, img); err != nil {
		log.Fatal(err)
	}
}
