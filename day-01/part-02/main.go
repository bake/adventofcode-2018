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
	s := bufio.NewScanner(f)
	changes := []int{}
	n := 0
	for s.Scan() {
		if _, err := fmt.Sscanf(s.Text(), "%d", &n); err != nil {
			log.Fatal(err)
		}
		changes = append(changes, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	freqs := map[int]bool{}
	for i := 0; true; i++ {
		i %= len(changes)
		sum += changes[i%len(changes)]
		if freqs[sum] {
			break
		}
		freqs[sum] = true
	}
	fmt.Println(sum)
}
