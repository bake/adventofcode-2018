package grid_test

import (
	"testing"

	"github.com/bakerolls/adventofcode-2018/day-11/grid"
)

func TestCellPower(t *testing.T) {
	tt := []struct {
		grid  grid.Grid
		x, y  int
		power int
	}{
		{8, 3, 5, 4},
		{57, 122, 79, -5},
		{39, 217, 196, 0},
		{71, 101, 153, 4},
	}
	for _, tc := range tt {
		power := tc.grid.Power(tc.x, tc.y)
		if power != tc.power {
			t.Fatalf("expected power grid with serial %d to be %d, got %d", tc.grid, tc.power, power)
		}
	}
}
