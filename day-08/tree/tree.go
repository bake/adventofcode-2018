package tree

// Node is the base of a tree.
type Node struct {
	Children []*Node
	Metadata []int
}

// New returns a root node and the remaining fields.
func New(fields []int) (*Node, []int, error) {
	root := &Node{}
	numChildren, numMetadata := fields[0], fields[1]
	fields = fields[2:]
	for i := 0; i < numChildren; i++ {
		node, remainder, err := New(fields)
		if err != nil {
			return nil, fields, err
		}
		fields = remainder
		root.Children = append(root.Children, node)
	}
	for i := 0; i < numMetadata; i++ {
		root.Metadata = append(root.Metadata, fields[0])
		fields = fields[1:]
	}
	return root, fields, nil
}

// Metasum returns the sum of the node and all its childrens metadata.
func (n *Node) Metasum() int {
	var sum int
	for _, m := range n.Metadata {
		sum += m
	}
	for _, n := range n.Children {
		sum += n.Metasum()
	}
	return sum
}

// Value calculates the value of a node. This is done by summing its metadata
// if it has no child nodes. If a node has children, its metadata refers to
// ther idices and Value() summs their metadata.
func (n *Node) Value() int {
	if len(n.Children) == 0 {
		return n.Metasum()
	}

	var sum int
	for _, m := range n.Metadata {
		if m-1 < 0 || m-1 >= len(n.Children) {
			continue
		}
		sum += n.Children[m-1].Value()
	}
	return sum
}
