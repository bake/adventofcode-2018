package pots

import "fmt"

type Pot bool

func (p Pot) String() string {
	if p {
		return "#"
	}
	return "."
}

type Pots []Pot

func (ps Pots) String() string {
	var str string
	for _, p := range ps {
		str += p.String()
	}
	return str
}

func (ps Pots) Match(i int, rule Pots) bool {
	for j := -2; j < 3; j++ {
		p := Pot(false)
		if i+j >= 0 && i+j < len(ps) {
			p = ps[i+j]
		}
		if p != rule[j+2] {
			return false
		}
	}
	return true
}

func (ps Pots) Evolve(rs []Rule) Pots {
	next := make(Pots, len(ps))
	for i := range next {
		for _, r := range rs {
			if ps.Match(i, r.Pots) {
				next[i] = r.Result
				break
			}
		}
	}
	return next
}

func (ps Pots) Equal(qs Pots) bool {
	for i := range qs {
		if ps[i] != qs[i] {
			return false
		}
	}
	return true
}

type Rule struct {
	Pots   Pots
	Result Pot
}

func (r Rule) String() string {
	return fmt.Sprintf("%s => %s", r.Pots, r.Result)
}

func Parse(str string) Pots {
	var ps Pots
	for _, r := range str {
		switch r {
		case '.':
			ps = append(ps, false)
		case '#':
			ps = append(ps, true)
		}
	}
	return ps
}
