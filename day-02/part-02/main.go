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
	var vals []string
	for s.Scan() {
		vals = append(vals, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for _, a := range vals {
		for _, b := range vals {
			n, str := diff(a, b)
			if n != 1 {
				continue
			}
			fmt.Println(str)
			return
		}
	}
}

func diff(a, b string) (int, string) {
	var c string
	var diffs int
	for i := range a {
		if a[i] == b[i] {
			c += string(a[i])
		}
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs, c
}
