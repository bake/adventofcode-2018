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
	fmt.Println(distance.Max(bs))
}
