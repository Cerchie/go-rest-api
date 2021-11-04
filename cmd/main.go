package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Cerchie/go-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
	router.HandleFunc("/haikus", handlers.GetAllHaikus).Methods(http.MethodGet)
	router.HandleFunc("/haikus/{id}", handlers.GetBook).Methods(http.MethodGet)
    router.HandleFunc("/haikus", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode("Hello World")
    })

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}