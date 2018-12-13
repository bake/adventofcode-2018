package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bakerolls/adventofcode-2018/day-13/mine"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	t := mine.Parse(bytes.Split(b, []byte("\n")))
	for t.Alive() > 1 {
		for t.Tick() {
		}
	}

	var x, y int
	for _, c := range t.Carts {
		if !c.Crashed {
			x, y = c.X, c.Y
			break
		}
	}
	fmt.Printf("%d,%d\n", x, y)
}
