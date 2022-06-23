package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const PORT = "8000"

func Serve() {
	mux := mux.NewRouter()

	mid := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}).Handler(mux)

	log.Println("Connected with Database")

	log.Fatal(http.ListenAndServe(":"+PORT, mid))

}
