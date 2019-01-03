package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/bakerolls/adventofcode-2018/day-18/grid"
)

func main() {
	raw, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	g := grid.Grid(raw)
	for i := 1; i <= 10; i++ {
		// fmt.Println(g)
		// if err := encode(fmt.Sprintf("out/%d.png", i), g.Image()); err != nil {
		// 	log.Fatal(err)
		// }
		g.Tick()
	}

	var trees, lumberyards int
	for _, b := range g {
		switch b {
		case grid.Trees:
			trees++
		case grid.Lumberyard:
			lumberyards++
		}
	}
	fmt.Println(trees * lumberyards)
}

func encode(name string, img image.Image) error {
	w, err := os.Create(name)
	defer w.Close()
	if err != nil {
		return err
	}
	if err := png.Encode(w, img); err != nil {
		return err
	}
	return nil
}
