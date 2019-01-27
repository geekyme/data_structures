package btree

import (
	"testing"
)

func TestNodeAddKey(t *testing.T) {
	node := Node{}

	node.AddKey(4, 0)

	if node.keys[0] != 0 {
		t.Errorf("Unexpected values of node %v", node)
	}

	node.AddKey(4, 2)

	if node.keys[0] != 0 || node.keys[1] != 2 {
		t.Errorf("Unexpected values of node %v", node)
	}

	node.AddKey(4, 1)

	if node.keys[0] != 0 || node.keys[1] != 1 || node.keys[2] != 2 {
		t.Errorf("Unexpected values of node %v", node)
	}
}
