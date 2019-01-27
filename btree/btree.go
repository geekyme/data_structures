package btree

/**
 * Order M tree:
 * Each node will have M children
 * Each node will have M-1 keys
 */

// Value represents the value to be stored in a node
type Value int

// Node contains values and pointers to children Nodes
type Node struct {
	keys     []Value
	children []*Node
}

// Tree has an order and a pointer to root Node
type Tree struct {
	order Value
	root  *Node
}

// Create a tree of order m
// func Create(order int) *Tree {
// 	t := Tree{order: order}
// 	return &t
// }

// func (t *Tree) Insert(value Value) *Tree {
// 	root := t.root
// }

// AddKey will add a value to the keys provided there is space
func (n *Node) AddKey(order int, value Value) *Node {
	l := len(n.keys)

	result := []Value{}

	var i int

	for i = 0; i < l; i++ {
		if el := n.keys[i]; value <= el {
			break
		} else {
			result = append(result, el)
		}
	}

	result = append(result, value)
	result = append(result, n.keys[i:]...)

	n.keys = result

	return n
}

// func (n *Node) Promote(order int) *Node {

// }
