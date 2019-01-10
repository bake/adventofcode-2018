package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bakerolls/adventofcode-2018/day-20/distance"
)

func main() {
	bs, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var total int
	for _, dist := range distance.All(bs) {
		if dist >= 1000 {
			total++
		}
	}
	fmt.Println(total)
}
