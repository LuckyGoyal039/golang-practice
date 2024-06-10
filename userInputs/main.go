package main

import "fmt"

func main() {

	var userName string
	fmt.Println("Please Enter your name: ")

	fmt.Scan(&userName)

	fmt.Println("Hello ", userName)

}
