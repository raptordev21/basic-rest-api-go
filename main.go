package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raptordev21/basic-rest-api-go/models"
	"github.com/raptordev21/basic-rest-api-go/routes"
)

func main() {
	r := mux.NewRouter()

	routes.NetflixRouter(r)
	routes.CourseRouter(r)

	models.Courses = append(models.Courses, models.Course{
		CourseId:    "1",
		CourseName:  "React",
		CoursePrice: 299,
		Author:      &models.Author{Fullname: "Ace", Website: "learnreact.com"},
	})
	models.Courses = append(models.Courses, models.Course{
		CourseId:    "2",
		CourseName:  "Go",
		CoursePrice: 199,
		Author:      &models.Author{Fullname: "Ace", Website: "learngo.com"},
	})

	r.HandleFunc("/", serveHome).Methods("GET")

	fmt.Println("Server running on PORT: 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in GO mux</h1>"))
}
