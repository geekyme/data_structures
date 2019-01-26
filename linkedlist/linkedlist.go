package linkedlist

// Post contains value of itself and reference to the next Post
type Post struct {
	value string
	next  *Post
}

// Feed is a collection of posts
type Feed struct {
	length int
	start  *Post
	end    *Post
}

// Append adds a new post to the feed
func (f *Feed) Append(newPost *Post) *Feed {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++

	return f
}
