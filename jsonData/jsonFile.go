package jsonFile

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

	// EncodingJson()

	DecodingJson()
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

func DecodingJson() {
	jsonData := []byte(`
	 {
                "Name": "react tutorials",
                "Price": 399,
                "Platform": "learnonline.com",
                "tags": [
                        "golang",
                        "frontend"
                ]
        }`)

	var lcoCourse course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// just want to addd data in key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
