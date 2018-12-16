package game_test

import (
	"bytes"
	"testing"

	"github.com/bakerolls/adventofcode-2018/day-15/game"
	"github.com/bakerolls/adventofcode-2018/day-15/pathfinding"
)

func TestTurn(t *testing.T) {
	tt := []struct {
		grid      string
		turns, hp int
	}{
		{
			grid:  "#######\n#.G...#\n#...EG#\n#.#.#G#\n#..G#E#\n#.....#\n#######",
			turns: 47,
			hp:    590,
		},
		{
			grid:  "#######\n#G..#E#\n#E#E.E#\n#G.##.#\n#...#E#\n#...E.#\n#######",
			turns: 37,
			hp:    982,
		},
		{
			grid:  "#######\n#E..EG#\n#.#G.E#\n#E.##E#\n#G..#.#\n#..E#.#\n#######",
			turns: 46,
			hp:    859,
		},
		{
			grid:  "#######\n#E.G#.#\n#.#G..#\n#G.#.G#\n#G..#.#\n#...E.#\n#######",
			turns: 35,
			hp:    793,
		},
		{
			grid:  "#######\n#.E...#\n#.#..G#\n#.###.#\n#E#G#G#\n#...#G#\n#######",
			turns: 54,
			hp:    536,
		},
		{
			grid:  "#########\n#G......#\n#.E.#...#\n#..##..G#\n#...##..#\n#...#...#\n#.G...G.#\n#.....G.#\n#########",
			turns: 20,
			hp:    937,
		},
	}

	for i, tc := range tt {
		grid := bytes.Split([]byte(tc.grid), []byte("\n"))
		g := game.New(grid, pathfinding.BreadthFirstSearch)
		for g.Turn() {
		}
		turns, hp := g.Outcome()
		if turns != tc.turns || hp != tc.hp {
			t.Fatalf("expected outcome of game %d to be (%d,%d), got (%d,%d)", i, tc.turns, tc.hp, turns, hp)
		}
	}
}
