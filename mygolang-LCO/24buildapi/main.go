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
	CourseName  string `json:"coursename"`
	CourseId    string `json:"courseid"`
	CoursePrice int    `json:"price"`
	Author      *Author
}

type Author struct {
	AuthorName    string `json:"authorname"`
	AuthorWebsite string `json:"website"`
}

//fake db

var courses []Course

//middleware, helper

func (c *Course) IsEmpty() bool {
	//return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

//controllers

func main() {
	fmt.Println("Buliding my first API using Golang")
	//var a bool:= IsEmpty()
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{AuthorName: "Satyam", AuthorWebsite: "www.lco.in"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "JavaScript", CoursePrice: 399,
		Author: &Author{AuthorName: "DubeyJi", AuthorWebsite: "www.lco.in"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCoursebyId).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listenig at port

	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Building thr Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing all available courses")

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(courses)
	// this process is AKA seeding, when we test our API, we seed some fake data to test
}

func getOneCoursebyId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing one course by ID")
	w.Header().Set("Content-type", "application/json")
	// grab id from request

	params := mux.Vars(r)
	fmt.Printf("Datatype of params: %T\n", params)

	// lopp through course and find matching by id
	for _, j := range courses {
		if j.CourseId == params["id"] {
			json.NewEncoder(w).Encode(j)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given id ")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing one course by ID")
	w.Header().Set("Content-type", "application/json")

	//what if, body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what if, data {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the Json")
		return
	}
	//generate unique id and conv them string
	//append new course in courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

//update one course

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating given course")
	w.Header().Set("Content-type", "application/json")
	// first grab id from req
	params := mux.Vars(r)

	//loop, find id, delete, add with my id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("Id not found")

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleted given course")
	w.Header().Set("Content-type", "application/json")

	var params = mux.Vars(r) // for grabbing id
	//var count = 0
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted")
			break
		}

	}
	json.NewEncoder(w).Encode("Course Id not found to delete ")

}
