package blogpost_test

import (
	"bytes"
	"testing"

	blogpost "github.com/zeufack/blogposts"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogpost.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tag:         []string{"go", "tdd"},
		}
	)

	t.Run("it cnverts a single post inot HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogpost.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1><p>This is a description</p>Tag: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
