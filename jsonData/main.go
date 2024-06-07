package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:coursename`
	Price    int
	Platform string
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Play with Json data")

	EncodingJson()
}

func EncodingJson() {
	myCourses := []course{
		{"Golang tutorials", 599, "learnonline.com", "abc@123", []string{"golang", "backend"}},
		{"react tutorials", 399, "learnonline.com", "abc@123", []string{"golang", "frontend"}},
		{"javascript tutorials", 99, "learnonline.com", "abc@123", nil},
	}

	// package this data as json data

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
