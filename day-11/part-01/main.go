package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
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

	width, height, size := 300, 300, 3

	g := grid.New(width, height, serial)
	x, y, max := g.Max(size)
	fmt.Printf("%d,%d\n", x, y)

	img := g.Image(width, height, max)
	w, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(w, img); err != nil {
		log.Fatal(err)
	}
}
