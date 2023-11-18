package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	req, err := http.Get("https://my-json-server.typicode.com/typicode/demo/posts")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	var posts []Post

	if err := json.NewDecoder(req.Body).Decode(&posts); err != nil {
		panic(err)
	}

	for _, post := range posts {
		fmt.Printf("ID: %d\nTitle: %s\n\n", post.ID, post.Title)
	}
}
