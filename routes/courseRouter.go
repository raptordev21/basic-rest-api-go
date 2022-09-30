package routes

import (
	"github.com/gorilla/mux"
	"github.com/raptordev21/basic-rest-api-go/controllers"
)

func CourseRouter(router *mux.Router) {
	router.HandleFunc("/courses", controllers.GetAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", controllers.GetOneCourse).Methods("GET")
	router.HandleFunc("/course", controllers.CreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}", controllers.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", controllers.DeleteCourse).Methods("DELETE")
}
