package distance_test

import (
	"testing"

	"github.com/bakerolls/adventofcode-2018/day-20/distance"
)

func TestMax(t *testing.T) {
	tt := []struct {
		regex []byte
		dist  int
	}{
		{[]byte("^WNE$"), 3},
		{[]byte("^ENWWW(NEEE|SSE(EE|N))$"), 10},
		{[]byte("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"), 18},
		{[]byte("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"), 23},
		{[]byte("^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"), 31},
	}
	for i, tc := range tt {
		if dist := distance.Max(tc.regex); dist != tc.dist {
			t.Fatalf("expected %d dist in test %d, got %d", tc.dist, i, dist)
		}
	}
}
