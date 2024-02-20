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

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // Author struct has not been declared yet. Hence using a pointer to it as Type.
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

func isEmpty(c *Course) bool {
	// return c.CourseId != "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline")

	r := mux.NewRouter()

	// seeding the database
	courses = append(
		courses,
		Course{
			CourseId:    "131",
			CourseName:  "Golang course",
			CoursePrice: 21,
			Author:      &Author{FullName: "Raj Karmakar", Website: "www.github.com"}})
	courses = append(
		courses,
		Course{
			CourseId:    "121",
			CourseName:  "React course",
			CoursePrice: 2113,
			Author:      &Author{FullName: "Samuel Mark", Website: "www.github1.com"}})

	// set router Controllers
	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a Port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers
// Server home
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Homepage. This server has been built using Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course has been found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil { // If req body is empty
		json.NewEncoder(w).Encode("Can not have empty body")
		return
	}

	var courseReq Course
	_ = json.NewDecoder(r.Body).Decode(&courseReq)
	if isEmpty(&courseReq) {
		json.NewEncoder(w).Encode("Can not have empty course body")
		return
	}

	// check if CourseName is a duplicate
	for _, course := range courses {
		if course.CourseName == courseReq.CourseName {
			json.NewEncoder(w).Encode("Course with same 'coursename' is already present")
			return
		}
	}

	// Generate unique courseId & convert it into a String
	// Append newly created course into database
	rand.Seed(time.Now().UnixNano())
	courseId := rand.Intn(100)
	courseReq.CourseId = strconv.Itoa(courseId)
	fmt.Println(courseId)

	courses = append(courses, courseReq)

	json.NewEncoder(w).Encode(courseReq)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req
	params := mux.Vars(r)
	// TODO: Send a response if id is not found

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			break
		}
	}

	json.NewEncoder(w).Encode(courses)
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(courses)
	return
}
