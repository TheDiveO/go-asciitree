package asciitree_test

import (
	"fmt"

	asciitree "github.com/thediveo/go-asciitree"
)

func Example_multiRoot() {
	// user-defined tree data structure with asciitree-related field tags.
	type node struct {
		Label    string   `asciitree:"label"`
		Props    []string `asciitree:"properties"`
		Children []node   `asciitree:"children"`
	}
	// set up two root nodes with their own tree of child nodes
	roots := []node{
		{
			Label: "root 1",
			Children: []node{
				{Label: "child 1", Props: []string{"childish"}},
				{Label: "child 2", Children: []node{
					{Label: "grandchild 1", Props: []string{"very childish"}},
					{Label: "grandchild 2"},
				}},
				{Label: "child 3"},
			},
		},
		{
			Label: "root 2",
			Children: []node{
				{Label: "child 2-1"},
			},
		},
	}
	// render the trees into a string and print it.
	fmt.Println(asciitree.RenderFancy(roots))
	// Output:
	// root 1
	// ├─ child 1
	// │     • childish
	// ├─ child 2
	// │  ├─ grandchild 1
	// │  │     • very childish
	// │  └─ grandchild 2
	// └─ child 3
	// root 2
	// └─ child 2-1
}
