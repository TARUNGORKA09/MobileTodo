package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TARUNGORKA09/MobileTodo/handlers"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server is starting ................")

	l := log.New(os.Stdout, "Mobile Todo", log.LstdFlags)

	mobile := handlers.NewMobile(l)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/getMobile/{id:[0-9]+}", mobile.GetMobileInfo)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/addMobile", mobile.AddMobile)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/deleteMobile/{id:[0-9]+}",mobile.DeleteMobile)

	http.ListenAndServe(":8080", sm)

}
