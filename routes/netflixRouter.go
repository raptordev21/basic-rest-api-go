package routes

import (
	"github.com/gorilla/mux"
	"github.com/raptordev21/basic-rest-api-go/controllers"
)

func NetflixRouter(router *mux.Router) {
	router.HandleFunc("/api/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/delete-all-movies", controllers.RemoveAllMovies).Methods("DELETE")
}

// func Router() *mux.Router {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/api/movies", controllers.GetMovies).Methods("GET")
// 	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
// 	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
// 	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
// 	router.HandleFunc("/api/delete-all-movies", controllers.RemoveAllMovies).Methods("DELETE")

// 	return router
// }
