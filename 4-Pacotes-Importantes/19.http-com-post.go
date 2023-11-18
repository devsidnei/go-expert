package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"Sidnei", "age":"undefined"}`))
	res, err := c.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)

}
