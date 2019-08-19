package input

import (
	df "../common"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type ServiceLoader struct {}

func (service ServiceLoader) GetPost() ([]df.Post, error) {
	resp, err := http.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []df.Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
func (service ServiceLoader) GetComment() ([]df.Comment, error) {
	resp, err := http.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []df.Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
