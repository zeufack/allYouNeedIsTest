package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tag         []string
	Body        string
}

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}
	var posts []Post

	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	// postData, err := io.ReadAll(postFile)
	// if err != nil {
	// 	return Post{}, err
	// }

	// post := Post{Title: string(postData)[7:]}

	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
	bodySeparator        = "---"
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// scanner.Scan()
	// postTitle := readLine(titleSeparator)

	// scanner.Scan()
	// postDesription := readLine(descriptionSeparator)

	// postData, err := io.ReadAll(postFile)

	// if err != nil {
	// 	return Post{}, err
	// }
	return Post{
		Title:       readLine(titleSeparator),
		Description: readLine(descriptionSeparator),
		Tag:         strings.Split(readLine(tagSeparator), ", "),
		Body:        readBody(scanner)}, nil
}
func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tag: <ul>{{range .Tag}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(buf io.Writer, aPost Post) error {

	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(buf, aPost); err != nil {
		return err
	}
	return err
}
