package tree

import (
	"testing"
)

//     1
//       \
//        5
//       / \
//      3   6
//     /
//    2
func setupTree() *Tree {
	nums := []int{1, 5, 6, 3, 2}

	tree := CreateTree(nums)

	return tree
}

func TestDFS(t *testing.T) {
	tree := setupTree()

	if DFS(tree.root) != "15326" {
		t.Errorf("Unexpected output: %v", DFS(tree.root))
	}
}

func TestBFS(t *testing.T) {
	tree := setupTree()

	if BFS(tree.root) != "15362" {
		t.Errorf("Unexpected output: %v", BFS(tree.root))
	}
}
