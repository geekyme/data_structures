package linkedlist

// Post contains value of itself and reference to the next Post
type Post struct {
	value string
	next  *Post
}

// Feed is a collection of posts
type Feed struct {
	start *Post
	end   *Post
}

// Append adds a new post to the feed
func (f *Feed) Append(p *Post) *Feed {
	if f.start == nil && f.end == nil {
		f.start = p
		f.end = p
	} else {
		current := f.start
		for current.next != nil {
			current = current.next
		}

		current.next = p
		f.end = p
	}

	return f
}
