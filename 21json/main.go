package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Encoding in Go")
	encodeJson()
	fmt.Println("JSON Encoding Completed Successfully!")
}

func encodeJson() {
	listCourse := []course{
		{"Anas", 120, "Udemy", "1234", []string{"Go", "Programming", "Backend"}},
		{"Ramesh", 420, "Youtube", "1240", []string{"React", "Programming", "Backend"}},
		{"Jitesh", 520, "Twitch", "5240", nil},
	}
	// finalJson, err := json.Marshal(listCourse)
	finalJson, err := json.MarshalIndent(listCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
