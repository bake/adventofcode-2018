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
	for t.Tick() {
		// Plot the map on every tick
		// fmt.Println(t)
	}
	x, y, _ := t.Crash()
	fmt.Printf("%d,%d", x, y)
}
