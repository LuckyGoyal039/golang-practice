package main

import "fmt"

func main() {

	var userName string
	var myarr [50]string
	var size int
	fmt.Println("Please Enter your name: ")

	fmt.Scan(&userName)

	fmt.Println("Please Enter Size of Array: ")

	fmt.Scan(&size)

	myarr[size] = "myName"

	fmt.Println("Hello ", userName)

	fmt.Println("array: ", myarr)
}
