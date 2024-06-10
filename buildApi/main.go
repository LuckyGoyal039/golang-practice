package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CouserId    string  `json:"courseid"`
	CourseName  string  `json: "coursename"`
	CoursePrice int     `json: "price"`
	Author      *Author `json: "author"`
}

type Author struct {
	Fullname string `jsong:"fullname"`
	Website  string `json :"website"`
}

// db
var courses []Course

// middleware, helper-files

func (c *Course) IsEmpty() bool {
	return c.CouserId == "" && c.CourseName == ""
}

func main() {

}

// controllers -file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang learning</h1>"))

}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
