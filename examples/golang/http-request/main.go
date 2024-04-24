package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return resBody, nil
}

func main() {
	res, err := Get("http://localhost:8080/")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
