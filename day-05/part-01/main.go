package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	polymer := string(b)

	for {
		p := react(polymer)
		if len(p) == 0 || len(p) == len(polymer) {
			break
		}
		polymer = p
	}
	fmt.Println(len(polymer))
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
