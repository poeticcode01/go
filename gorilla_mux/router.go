package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck)
	fmt.Println("Server is getting up...")
	http.ListenAndServe(":8080", router)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"msg": "Server is up and running" }`)
}
