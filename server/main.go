package main

import (
	"log"
	"net/http"

	. "github.com/JuanJoseGonGi/recetas-restapi/receta"
	"github.com/gorilla/handlers"
)

func main() {
	r := NewRouter()
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))
}
