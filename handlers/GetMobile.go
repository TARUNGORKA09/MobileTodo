package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type mobile struct {
	l log.Logger
}

func getMobile(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(rw, "Unable to convert")
	}
	mob := data.getMobile(id)
	err1 := mob.toJSON(rw)
	if err1 != nil {
		fmt.Printf("unable to convert to JSON")
	}

}
