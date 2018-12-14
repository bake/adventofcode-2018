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

// Suffix tests if a slice of recipe scores ends with another.
func (rs Recipes) Suffix(es Recipes) bool {
	if len(rs) < len(es) {
		return false
	}
	for i, r := range rs[len(rs)-len(es):] {
		if r != es[i] {
			return false
		}
	}
	return true
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var target Recipes
	for _, b := range strings.TrimSpace(string(b)) {
		t, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatal(err)
		}
		target = append(target, uint8(t))
	}

	elfs := []int{0, 1}
	recipes := Recipes{3, 7}
	for {
		var score uint8
		for _, elf := range elfs {
			score += recipes[elf]
		}

		if score >= 10 {
			recipes = append(recipes, score/10)
			if recipes.Suffix(target) {
				fmt.Println(len(recipes) - len(target))
				break
			}
		}
		recipes = append(recipes, score%10)
		if recipes.Suffix(target) {
			fmt.Println(len(recipes) - len(target))
			break
		}

		for i, elf := range elfs {
			elfs[i] = (elf + 1 + int(recipes[elf])) % len(recipes)
		}
	}
}
