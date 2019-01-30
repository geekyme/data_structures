package tree

import "strconv"

// Node represents the nodes in the tree, with a binary structure
type Node struct {
	left  *Node
	right *Node
	value int
}

// Tree contains a pointer to the root Node
type Tree struct {
	root *Node
}

// CreateTree creates a binary search tree from a slice of numbers
func CreateTree(nums []int) *Tree {
	t := &Tree{}

	for _, num := range nums {
		newNode := &Node{value: num}
		if t.root == nil {
			t.root = newNode
		} else {
			insert(t.root, newNode)
		}
	}

	return t
}

func insert(node, newNode *Node) {

	if node.value > newNode.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insert(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insert(node.right, newNode)
		}
	}
}

// DFS performs a depth-first search on the node and output the sequence of values
func DFS(n *Node) string {
	stack := []*Node{}

	stack = append(stack, n)

	res := ""

	for len(stack) != 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[0:last]

		res = res + strconv.Itoa(node.value)

		if node.right != nil {
			stack = append(stack, node.right)
		}

		if node.left != nil {
			stack = append(stack, node.left)
		}
	}

	return res
}

// BFS performs a breadth-first search on the node and output the sequence of values
func BFS(n *Node) string {
	queue := []*Node{}

	queue = append(queue, n)

	res := ""

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		res = res + strconv.Itoa(node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return res
}
