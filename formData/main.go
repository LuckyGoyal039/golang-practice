package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Play with form data")

	PostFormData()

}

func PostFormData() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	var data = url.Values{}

	data.Add("firstName", "Lucky")
	data.Add("lastName", "Goyal")
	data.Add("email", "Lucky@go.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	defer response.Body.Close()
	// one way
	fmt.Println("my Content is: ", string(content))

	// anothe way
	var responseString strings.Builder

	byteCode, _ := responseString.Write(content)

	fmt.Println("ByteCode is: ", byteCode)

	fmt.Println("Content is: ", responseString.String())
}
