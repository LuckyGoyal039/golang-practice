package main

import "fmt"

func main() {
	user1 := User{"Lucky Goyal", 24}

	user1.showDetails()
}

type User struct {
	Name string
	Age  int16
}

func (u User) showDetails() {
	fmt.Printf("my name is: %v and my age is: %v", u.Name, u.Age)
}
