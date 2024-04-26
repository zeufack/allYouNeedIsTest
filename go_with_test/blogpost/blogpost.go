package blogpost

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
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

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	postTitle := scanner.Text()

	scanner.Scan()
	postDesription := scanner.Text()

	// postData, err := io.ReadAll(postFile)

	// if err != nil {
	// 	return Post{}, err
	// }
	return Post{Title: string(postTitle)[7:], Description: string(postDesription)[13:]}, nil
}
