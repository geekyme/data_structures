package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := Queue{}

	item1 := &Item{value: 1}
	item2 := &Item{value: 2}
	item3 := &Item{value: 3}

	queue.Push(item1)
	queue.Push(item2)
	queue.Push(item3)

	if queue.size != 3 {
		t.Errorf("Unexpected size: %v", queue.size)
	}

	if item := queue.Pop(); item != item1 {
		t.Errorf("Unexpected item popped from queue: %v", item)
	}

	if item := queue.Pop(); item != item2 {
		t.Errorf("Unexpected item popped from queue: %v", item)
	}

	if item := queue.Pop(); item != item3 {
		t.Errorf("Unexpected item popped from queue: %v", item)
	}

	if queue.size != 0 {
		t.Errorf("Unexpected size: %v", queue.size)
	}

	if item := queue.Pop(); item != nil {
		t.Errorf("Unexpected item popped from queue: %v", item)
	}

	if queue.size != 0 {
		t.Errorf("Unexpected size: %v", queue.size)
	}
}
