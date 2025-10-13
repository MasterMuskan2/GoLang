package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for Course - File

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//Database

var courses []Course

//Middle-ware or, helper - file

func (c *Course) isEmpty() bool {
	// return (c.CourseId == "" && c.CourseName == "")
	return c.CourseName == ""
}

func main() {

}

//Controllers - file

//serve home route

func serveHomeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API Website</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All the Courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, value := range courses {
		if value.CourseId == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")

	// Body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide some data!")
	}

	// Body is {}

	var course Course

	value = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data found!")
		return
	}

	// Generate unqiue id as int then convert it to string and insert in db

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update the course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses, course)
			json.NewEncoder(w).Encode(course)
			return 
		}
	}
}
