package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type xy struct{ x, y int }
type claim struct{ id, x, y, w, h int }

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fabric := map[xy][]claim{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var c claim
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < c.w; i++ {
			for j := 0; j < c.h; j++ {
				xy := xy{c.x + i, c.y + j}
				fabric[xy] = append(fabric[xy], c)
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for _, cs := range fabric {
		if len(cs) != 1 {
			continue
		}
		if !overlap(cs[0], fabric) {
			fmt.Println(cs[0].id)
			return
		}
	}
}

func overlap(c claim, fabric map[xy][]claim) bool {
	for i := 0; i < c.w; i++ {
		for j := 0; j < c.h; j++ {
			xy := xy{c.x + i, c.y + j}
			if len(fabric[xy]) != 1 {
				return true
			}
		}
	}
	return false
}
