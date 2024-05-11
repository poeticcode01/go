package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/poeticcode01/go/web/routes"
)

func main() {
	fmt.Println("Starting web server")
	router := mux.Router{}
	routes.RoutesHandler(&router)
	http.ListenAndServe(":8080", &router)

}
