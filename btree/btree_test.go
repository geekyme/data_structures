package btree

import (
	"testing"
)

func TestNodeAddKey(t *testing.T) {
	node := Node{}

	// Add 0
	if _, err := node.AddKey(4, 0); err != nil {
		t.Errorf("Unexpected failure to add key when node is not at max capacity")
	}

	if node.keys[0] != 0 {
		t.Errorf("Unexpected values of node %v", node)
	}

	// Add 2
	if _, err := node.AddKey(4, 2); err != nil {
		t.Errorf("Unexpected failure to add key when node is not at max capacity")
	}

	if node.keys[0] != 0 || node.keys[1] != 2 {
		t.Errorf("Unexpected values of node %v", node)
	}

	// Add 1
	if _, err := node.AddKey(4, 1); err != nil {
		t.Errorf("Unexpected failure to add key when node is not at max capacity")
	}

	if node.keys[0] != 0 || node.keys[1] != 1 || node.keys[2] != 2 {
		t.Errorf("Unexpected values of node %v", node)
	}

	if n, err := node.AddKey(4, 2); err == nil {
		t.Errorf("Unexpectedly added a key when node is supposed to be at max capacity: %v", n)
	}
}
