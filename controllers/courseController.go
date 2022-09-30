package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/raptordev21/basic-rest-api-go/models"
)

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range models.Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course models.Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	models.Courses = append(models.Courses, course)

	json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var courseData models.Course
	_ = json.NewDecoder(r.Body).Decode(&courseData)
	if courseData.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	params := mux.Vars(r)

	for index, course := range models.Courses {
		if course.CourseId == params["id"] {
			models.Courses = append(models.Courses[:index], models.Courses[index+1:]...)
			courseData.CourseId = params["id"]
			models.Courses = append(models.Courses, courseData)
			json.NewEncoder(w).Encode(courseData)
			return
		}
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range models.Courses {
		if course.CourseId == params["id"] {
			models.Courses = append(models.Courses[:index], models.Courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Deleted Successfully")
}
