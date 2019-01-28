package queue

// Item represents the item in the queue
type Item struct {
	value interface{}
}

// Queue holds a collection of items in a FIFO
type Queue struct {
	items []*Item
	size  int
}

// Push adds an item into the queue
func (q *Queue) Push(item *Item) {
	q.items = append(q.items, item)
	q.size++
}

// Pop removes the first item from the queue
func (q *Queue) Pop() *Item {

	if q.size > 0 {
		item, items := q.items[0], q.items[1:]

		q.items = items

		if q.items != nil {
			q.size--
		}

		return item
	}

	return nil
}
