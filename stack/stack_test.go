package stack

import "testing"

func TestStack(t *testing.T) {
	stack := Stack{}

	item1 := &Item{value: 1}
	item2 := &Item{value: 2}
	item3 := &Item{value: 3}

	stack.Push(item1)
	stack.Push(item2)
	stack.Push(item3)

	if stack.size != 3 {
		t.Errorf("Unexpected size: %v", stack.size)
	}

	if item := stack.Pop(); item != item3 {
		t.Errorf("Unexpected item popped from stack: %v", item)
	}

	if item := stack.Pop(); item != item2 {
		t.Errorf("Unexpected item popped from stack: %v", item)
	}

	if item := stack.Pop(); item != item1 {
		t.Errorf("Unexpected item popped from stack: %v", item)
	}

	if stack.size != 0 {
		t.Errorf("Unexpected size: %v", stack.size)
	}

	if item := stack.Pop(); item != nil {
		t.Errorf("Unexpected item popped from stack: %v", item)
	}

	if stack.size != 0 {
		t.Errorf("Unexpected size: %v", stack.size)
	}
}
