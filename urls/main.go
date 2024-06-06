package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/comments?postId=1"

func main() {
	fmt.Println("Handling URLs")

	fmt.Println(myurl)

	// get information from url
	result, _ := url.Parse(myurl)
	fmt.Println("url Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Path: ", result.Path)
	fmt.Println("Query parameters: ", result.RawQuery)

	qparams := result.Query()
	fmt.Println("post Id is: ", qparams["postId"])

	//construct url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "jsonplaceholder.typicode.com",
		Path:    "/comments",
		RawPath: "postId=3",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
