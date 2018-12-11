package node

// Node holds the names of its ancestors and descendants.
type Node struct {
	Name        rune
	Ancestors   map[rune]bool
	Descendants map[rune]bool
}

// New creates a new node.
func New(name rune) *Node {
	return &Node{name, map[rune]bool{}, map[rune]bool{}}
}

// Ready returns true if all ancestors are inside done.
func (n Node) Ready(done map[rune]bool) bool {
	for a := range n.Ancestors {
		if !done[a] {
			return false
		}
	}
	return true
}

// Strings returns a nodes name as a stirng.
func (n Node) String() string {
	return string(n.Name)
}

// Nodes implements the sort interface.
type Nodes []*Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].Name < n[j].Name }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (n Nodes) String() string {
	str := ""
	for _, n := range n {
		str += n.String()
	}
	return str
}
