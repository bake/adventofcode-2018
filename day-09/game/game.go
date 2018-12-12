package game

import (
	"fmt"
)

// Game implements the marble game logic.
type Game struct {
	round      int
	numPlayers int
	numMarbles int

	Current *Node
	Players map[int]int
}

// New returns a new game.
func New(players, marbles int) *Game {
	g := &Game{
		numPlayers: players,
		numMarbles: marbles,
		Players:    map[int]int{},
		Current:    &Node{Value: 0},
	}
	g.Current.Prev, g.Current.Next = g.Current, g.Current
	return g
}

// Round starts the next players turn. Returns true until done.
func (g *Game) Round() bool {
	g.round++
	if g.round > g.numMarbles {
		return false
	}

	// Every time an elf gets a marble that is a multiple of 23, this marble
	// and the 7th-last marble adds to their points.
	if g.round%23 == 0 {
		p := g.round % g.numPlayers
		g.Players[p] += g.round
		marble := g.Current
		for i := 0; i < 7; i++ {
			marble = marble.Prev
		}
		g.Players[p] += marble.Value
		marble.Prev.Next = marble.Next
		marble.Next.Prev = marble.Prev
		g.Current = marble.Next
		return true
	}

	// Replace the 2nd next marble.
	n := &Node{
		Prev:  g.Current.Next,
		Next:  g.Current.Next.Next,
		Value: g.round,
	}
	g.Current.Next.Next.Prev = n
	g.Current.Next.Next = n
	g.Current = n
	return true
}

// Node is part of a linked list.
type Node struct {
	Prev, Next *Node
	Value      int
}

// String returns a string containing a lists values. It stopps when the value
// of the starting node appears a second time.
func (n *Node) String() string {
	var str string
	value := n.Value
	for {
		str += fmt.Sprintf("%d ", n.Value)
		n = n.Next
		if n.Value == value {
			break
		}
	}
	return str
}
