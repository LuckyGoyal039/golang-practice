package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://jsonplaceholder.typicode.com/posts/1"

func main() {

	responce, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	// fmt.Println(responce)
	defer responce.Body.Close()
	databytes, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
