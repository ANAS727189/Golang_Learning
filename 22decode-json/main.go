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
	fmt.Println("Welcome to JSON Decoding in Go")
	decodeJson()
	fmt.Println("JSON Decoding Completed Successfully!")
}

func decodeJson() {
	jsonDataFromWeb := []byte(`  
		{"coursename": "Anas","Price": 120,"website": "Udemy","tags": ["Go","Programming","Backend"]
        }
		`)

	var copyCourseData course
	ValidCourse := json.Valid(jsonDataFromWeb)
	if ValidCourse {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &copyCourseData)
		fmt.Printf("%#v\n", copyCourseData)
	} else {
		fmt.Println("JSON is not valid")
	}
	// Some cases when you only want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v, Value: %v and Type: %T", k, v, v)
	}
}
