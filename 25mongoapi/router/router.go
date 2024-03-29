package router

import (
	"github.com/gorilla/mux"
	"github.com/raj5036/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/watched", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
