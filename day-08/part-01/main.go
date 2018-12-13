package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/bakerolls/adventofcode-2018/day-08/tree"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	body := strings.Fields(string(b))
	fields := make([]int, len(body))
	for i, b := range body {
		fields[i], err = strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}
	}

	n, _, err := tree.New(fields)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n.Metasum())
}
