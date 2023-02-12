package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

var Users []User

func RequestHandler(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", HealthCheck)
	router.HandleFunc("/create", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/count", CountAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetSingleUser).Methods("GET")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":10001", handlers.CORS(headers, methods, origins)(router)))
}

func main(){
	RequestHandler()
}

