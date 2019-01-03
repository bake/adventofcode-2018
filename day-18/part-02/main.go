package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bakerolls/adventofcode-2018/day-18/grid"
)

func main() {
	grids := map[string]int{}
	res := []int{}
	raw, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	g := grid.Grid(raw)

	// Similar to a previous problem, this one loops. It helps plotting each
	// step of part-01 to see it. We just need to iterate until we see a grid we
	// already know and then calculate the steps index that would be euqal to the
	// 1000000000th tick.
	max := 1000000000
	var start, end int
	for i := 0; i < max; i++ {
		if s, ok := grids[string(g)]; ok {
			start, end = s, i
			break
		}
		grids[string(g)] = i
		res = append(res, resources(g))
		g.Tick()
	}
	fmt.Println(res[start+(max-start)%(end-start)])
}

func resources(g grid.Grid) int {
	var trees, lumberyards int
	for _, b := range g {
		switch b {
		case grid.Trees:
			trees++
		case grid.Lumberyard:
			lumberyards++
		}
	}
	return trees * lumberyards
}
