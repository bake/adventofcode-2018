package tree_test

import (
	"testing"

	"github.com/bakerolls/adventofcode-2018/day-08/tree"
)

func TestParseNode(t *testing.T) {
	tt := []struct {
		name   string
		fields []int
		node   *tree.Node
	}{
		{
			"no children but metadata",
			[]int{0, 1, 99},
			&tree.Node{Children: []*tree.Node{}, Metadata: []int{99}},
		},
		{
			"children and metadata",
			[]int{1, 1, 0, 1, 99, 2},
			&tree.Node{
				Children: []*tree.Node{{Children: nil, Metadata: []int{99}}},
				Metadata: []int{2},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			node, _, err := tree.New(tc.fields)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}
			if len(node.Children) != len(tc.node.Children) {
				t.Fatalf("expected node to have %d children, got %d", len(tc.node.Children), len(node.Children))
			}
			if len(node.Metadata) != len(tc.node.Metadata) {
				t.Fatalf("expected metadata have %d fields, got %v", len(tc.node.Metadata), len(node.Metadata))
			}
			for i := range tc.node.Metadata {
				if node.Metadata[i] != tc.node.Metadata[i] {
					t.Fatalf("expected metadata %d to be %d, got %d", i, tc.node.Metadata[i], node.Metadata[i])
				}
			}
		})
	}
}
