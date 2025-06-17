package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("Hello, World!")
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseMessage{Message: "Hello, World!"})
	}).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
	log.Fatal((http.ListenAndServe(":8080", router)))
}
