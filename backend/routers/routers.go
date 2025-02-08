package routers

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	routers := mux.NewRouter()

	routers.HandleFunc("/list", handlers.GetAllList).Methods("GET")
	routers.HandleFunc("/list/{id}", handlers.GetList).Methods("GET")
	routers.HandleFunc("/list", handlers.CreateList).Methods("POST")
	routers.HandleFunc("/list/{id}", handlers.UpdateList).Methods("PUT")
	routers.HandleFunc("/list/{id}", handlers.DeleteList).Methods("DELETE")
	// Di main.go atau router setup
	routers.HandleFunc("/list/{id}/status", handlers.UpdateListStatus).Methods("PATCH")

	return routers
}
