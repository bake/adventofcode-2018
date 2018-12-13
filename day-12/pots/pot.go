package pots

import "fmt"

// Pot is true if a plant is growing inside it.
type Pot bool

// String returns a string representation of the plant.
func (p Pot) String() string {
	if p {
		return "#"
	}
	return "."
}

// Pots os a slie of pots.
type Pots []Pot

// String returns a string representation of the plants.
func (ps Pots) String() string {
	var str string
	for _, p := range ps {
		str += p.String()
	}
	return str
}

// Match returns true if a given rule matches the pots at position i.
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

// Evolve evolves the plants.
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

// Equal returns true if two pot slices are equal.
func (ps Pots) Equal(qs Pots) bool {
	for i := range qs {
		if ps[i] != qs[i] {
			return false
		}
	}
	return true
}

// Rule holds a rule and if a plant grows when the rule matches.
type Rule struct {
	Pots   Pots
	Result Pot
}

// String returns a string representation of the rule.
func (r Rule) String() string {
	return fmt.Sprintf("%s => %s", r.Pots, r.Result)
}

// Parse returns a slice of pots. Characters that are not . or # are ignored.
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
