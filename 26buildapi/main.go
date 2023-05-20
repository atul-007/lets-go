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
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {

	fmt.Println("build  api in go")
	r := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 200, Author: &Author{FullName: "Atul Ranjan", Website: "blabla.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "mern", CoursePrice: 200, Author: &Author{FullName: "Atul Ranjan", Website: "blabla.com"}})

	//routing
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getallcourses).Methods("GET")
	r.HandleFunc("/course/{id}", getonecourse).Methods("GET")
	r.HandleFunc("/course", createonecourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateonecourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteonecourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

func servehome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>API buid by Atul Ranjan</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	for _, Course := range courses {
		if Course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(Course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	//return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//when the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please enter some data ")
	}

	//when {}
	var course Course

	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside json")
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	//return
}

func updateonecourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	//first grab id from request
	params := mux.Vars(r)
	//loop to find the given id
	for index, course := range courses {

		if course.CourseId == params["id"] {

			courses = append(courses[:index], courses[index+1:]...) //delete the  old id that is to be updated
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) //decode the given data

			course.CourseId = params["id"]    //to make sure the id does not change cause its an update opertion
			courses = append(courses, course) //append the new data into the slice
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteonecourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("One course have been deleted")
			break
		}
	}
}
