package main

import (
	"backend/routers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := routers.SetupRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		Debug:          true,
	})

	handler := c.Handler(router)

	log.Println("Server sedang berjalan di port :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
