package node

import (
	"bufio"
	"fmt"
	"io"
)

// Parse parses nodes form an io.Reader.
func Parse(r io.Reader) (map[rune]*Node, error) {
	nodes := map[rune]*Node{}
	s := bufio.NewScanner(r)
	for i := 0; s.Scan(); i++ {
		var a, b rune
		_, err := fmt.Sscanf(s.Text(), "Step %c must be finished before step %c can begin.", &a, &b)
		if err != nil {
			return nil, err
		}
		if nodes[a] == nil {
			nodes[a] = New(a)
		}
		nodes[a].Descendants[b] = true
		if nodes[b] == nil {
			nodes[b] = New(b)
		}
		nodes[b].Ancestors[a] = true
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return nodes, nil
}
