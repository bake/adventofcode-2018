package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Recipes is a slice of recipe scores.
type Recipes []uint8

func (rs Recipes) String() string {
	var str string
	for _, r := range rs {
		str += strconv.Itoa(int(r))
	}
	return str
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		log.Fatal(err)
	}

	width := 10
	elfs := []int{0, 1}
	recipes := Recipes{3, 7}
	for i := 0; i < total+width; i++ {
		var score uint8
		for _, elf := range elfs {
			score += recipes[elf]
		}
		if score >= 10 {
			i++
			recipes = append(recipes, score/10)
		}
		recipes = append(recipes, score%10)
		for i, elf := range elfs {
			elfs[i] = (elf + 1 + int(recipes[elf])) % len(recipes)
		}
	}

	fmt.Println(recipes[total : total+width])
}
