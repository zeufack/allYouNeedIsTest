package blogpost_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogpost "github.com/zeufack/blogposts"
)

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPost(t *testing.T) {
	const (
		firstData = `Title: Post 1
Description: this is first data
Tags: tdd, go
---
Hello
World`
		secondData = `Title: Post 2
Description: this is second data
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstData)},
		"hello-world2.md": {Data: []byte(secondData)},
	}

	posts, err := blogpost.NewPostFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogpost.Post{
		Title:       "Post 1",
		Description: "this is first data",
		Tag:         []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v, want %v", got, want)
	// }

	assertPost(t, want, got)
}

func assertPost(t *testing.T, want, got blogpost.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v\n, want %+v\n", got, want)
	}
}
