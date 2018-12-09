package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"unicode"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	polymer := string(b)

	min := math.MaxInt32
	for _, u := range units(polymer) {
		newPoly := collapse(remove(polymer, u))
		fmt.Println(string(u), len(newPoly))
		if len(newPoly) < min {
			min = len(newPoly)
		}
	}
	fmt.Println(min)
}

func units(polymer string) []rune {
	unitMap := map[rune]bool{}
	for _, u := range polymer {
		unitMap[unicode.ToLower(u)] = true
	}
	var units []rune
	for u := range unitMap {
		units = append(units, u)
	}
	return units
}

func remove(polymer string, unit rune) string {
	var str string
	for _, r := range polymer {
		if unicode.ToLower(r) != unicode.ToLower(unit) {
			str += string(r)
		}
	}
	return str
}

func collapse(polymer string) string {
	for {
		p := react(polymer)
		if len(p) == 0 || len(p) == len(polymer) {
			break
		}
		polymer = p
	}
	return polymer
}

func react(polymer string) string {
	var str string
	var l rune
	for i, r := range polymer {
		if i == 0 {
			l = r
			continue
		}
		if unicode.ToLower(l) == unicode.ToLower(r) && unicode.IsLower(l) != unicode.IsLower(r) {
			return polymer[:i-1] + polymer[i+1:]
		}
		l = r
	}
	return str
}
