package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var twos, threes int
	s := bufio.NewScanner(f)
	for s.Scan() {
		nums := map[rune]int{}
		for _, r := range s.Text() {
			nums[r]++
		}
		var two, three int
		for _, n := range nums {
			if n == 2 {
				two = 1
			}
			if n == 3 {
				three = 1
			}
		}
		twos += two
		threes += three
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(twos * threes)
}

func count(str string) map[rune]int {
	nums := map[rune]int{}
	for _, r := range str {
		nums[r]++
	}
	return nums
}
