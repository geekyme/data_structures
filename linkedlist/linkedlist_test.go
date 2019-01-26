package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	post1 := Post{}
	post2 := Post{}
	feed := Feed{}

	feed.Append(&post1)

	if feed.start != &post1 || feed.end != &post1 {
		t.Errorf("Did not append properly into an empty feed: %s", feed)
	}

	feed.Append(&post2)

	if feed.start != &post1 {
		t.Errorf("Unexpectedly changed feed's first post: %s", feed)
	}

	if feed.end != &post2 {
		t.Errorf("Did not append post to the last: %s", feed)
	}
}
