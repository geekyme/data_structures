package stack

// Item represents the item in the stack
type Item struct {
	value interface{}
	next  *Item
}

// Stack holds a collection of items, LIFO style
type Stack struct {
	top  *Item
	size int
}

// Push will add a new item to the top of the stack
func (s *Stack) Push(item *Item) {
	if s.top != nil {
		prev := s.top
		s.top = item
		item.next = prev
	} else {
		s.top = item
	}

	s.size++
}

// Pop returns the item at the top of the stack
func (s *Stack) Pop() *Item {
	item := s.top

	if item != nil {
		s.top = item.next

		s.size--
	}

	return item
}
