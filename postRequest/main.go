package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Perform Post request")

	performPostRequest()
}

func performPostRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	requestBody := strings.NewReader(`{
	"userId":1,
	"id":1,
	"title":"This is my title"
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	defer response.Body.Close()

	// one way
	fmt.Println("content: ", string(content))

	// another way
	var responseString strings.Builder
	byteCode, _ := responseString.Write(content)

	fmt.Println("Bytecode is: ", byteCode)

	fmt.Println("My content is: ", responseString.String())
}
