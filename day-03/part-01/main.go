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

	type xy struct{ x, y int }
	fab := map[xy]int{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				fab[xy{x + i, y + j}]++
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	for _, n := range fab {
		if n > 1 {
			sum++
		}
	}
	fmt.Println(sum)
}
