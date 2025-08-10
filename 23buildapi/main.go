package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Encode = Go struct → JSON → send to frontend.
// Decode = JSON from frontend → Go struct → use in backend.

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Auther      *Auther `json:"auther"`
}

type Auther struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Creating fake database
var courses []Course

// middlewares
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Course API - Build Custom API with Golang")
	r := mux.NewRouter()
	//Seeding
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "ReactJS Bootcamp",
		CoursePrice: 299,
		Auther:      &Auther{Fullname: "John Doe", Website: "johndoe.com"},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "MERN Stack Bootcamp",
		CoursePrice: 399,
		Auther:      &Auther{Fullname: "Jane Doe", Website: "janedoe.com"},
	})
	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "Python Bootcamp",
		CoursePrice: 199,
		Auther:      &Auther{Fullname: "Alice Smith", Website: "alicesmith.com"},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "JavaScript Bootcamp",
		CoursePrice: 249,
		Auther:      &Auther{Fullname: "Bob Johnson", Website: "bobjohnson.com"},
	})
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start the server on port 8080
	fmt.Println("Server started on port 8080")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Course API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //Convert courses slice(which is a struct) to JSON and send it in response
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//Grab id from request
	params := mux.Vars(r)
	//Loop through all courses and find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}
	// Check if course name already exists
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course with this name already exists.")
			return
		}
	}
	//Generate a unique ID for the course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	fmt.Println("Course created successfully:", course.CourseName)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	// loop, id, remove that id, add my new course id
	w.Header().Set("Content-Type", "application/json")
	// Grab id from request
	params := mux.Vars(r)
	// Loop through all courses and find matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			var updatedCourse Course
			err := json.NewDecoder(r.Body).Decode(&updatedCourse)
			if err != nil {
				json.NewEncoder(w).Encode("Failed to update course")
				return
			}
			if updatedCourse.IsEmpty() {
				json.NewEncoder(w).Encode("Please provide course data")
				return
			}
			// courses = append(courses[:index], courses[index+1:]...) // Remove the course with the matching id
			// updatedCourse.CourseId = params["id"]                   // Assign the same id to the new course
			// courses = append(courses, updatedCourse)                // Add the new course to the slice
			//Directly updating the slice, instead of remove + append
			updatedCourse.CourseId = params["id"]
			courses[index] = updatedCourse // Update the course at the index
			json.NewEncoder(w).Encode(updatedCourse)
			fmt.Println("Course updated successfully:", updatedCourse.CourseName)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// Grab id from request
	params := mux.Vars(r)
	// Loop through all courses and find matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // Remove the course with the matching id
			json.NewEncoder(w).Encode(fmt.Sprintf("%v Course deleted successfully", course.CourseName))
			fmt.Println("Course deleted successfully:", course.CourseName)
			break
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")
	courses = []Course{} // Reset the courses slice to an empty slice
	json.NewEncoder(w).Encode("All courses deleted successfully")
	fmt.Println("All courses deleted successfully")
}
