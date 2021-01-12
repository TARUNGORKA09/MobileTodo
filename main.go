package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server is starting ................")

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getMobile/{id:[0-9]+}", handlers.getMobile)

}
