package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	post1 := Post{}
	post2 := Post{}
	feed := Feed{}

	feed.Append(&post1)

	if feed.start != &post1 || feed.end != &post1 {
		t.Errorf("Did not append properly into an empty feed: %v", feed)
	}

	feed.Append(&post2)

	if feed.start != &post1 {
		t.Errorf("Unexpectedly changed feed's first post: %v", feed)
	}

	if feed.end != &post2 {
		t.Errorf("Did not append post to the last: %v", feed)
	}
}

func BenchmarkAppend10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Append(10000)
	}
}

func BenchmarkAppend100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Append(100000)
	}
}

func Append(count int) {
	feed := Feed{}

	for n := 0; n < count; n++ {
		post := Post{}
		feed.Append(&post)
	}
}
