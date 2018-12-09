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
	var n, sum int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &n)
		sum += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
