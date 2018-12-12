package main

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/bakerolls/adventofcode-2018/day-10/canvas"
)

func main() {
	r, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var pts canvas.Points
	s := bufio.NewScanner(r)
	for s.Scan() {
		p := canvas.Point{}
		_, err := fmt.Sscanf(s.Text(), "position=<%d, %d> velocity=<%d, %d>", &p.X, &p.Y, &p.VX, &p.VY)
		if err != nil {
			log.Fatal(err)
		}
		pts = append(pts, p)
	}

	var width, height int
	for i := 0; ; i++ {
		for i, p := range pts {
			pts[i].X += p.VX
			pts[i].Y += p.VY
		}

		// Update until the canvas starts to extend again.
		b := pts.Bounds()
		newWidth, newHeight := b.Max.X-b.Min.X, b.Max.Y-b.Min.Y
		if i > 0 && newWidth >= width && newHeight >= height {
			break
		}
		width, height = newWidth, newHeight
	}
	for i, p := range pts {
		pts[i].X -= p.VX
		pts[i].Y -= p.VY
	}

	w, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(w, pts.Image()); err != nil {
		log.Fatal(err)
	}
}
