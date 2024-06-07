package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Learn how to send GET request")

	PerformGetRequest()

}

func PerformGetRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts/1"

	responce, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()
	fmt.Println("Status Code: ", responce.StatusCode)
	fmt.Println("Content length is: ", responce.ContentLength)

	content, _ := io.ReadAll(responce.Body)

	// one way to get the content
	fmt.Println(string(content))

	// another way
	var responceString strings.Builder

	byteCount, _ := responceString.Write(content)

	fmt.Println("Bytecode is: ", byteCount)

	fmt.Println("String content is: ", responceString.String())

}
