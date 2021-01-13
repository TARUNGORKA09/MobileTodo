package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/TARUNGORKA09/MobileTodo/data"
	"github.com/gorilla/mux"
)

func (m *MobileInfo) DeleteMobile(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["Id"])
	if err != nil {
		fmt.Fprintf(rw, "Unable to conevrt from string to int", id)
	}
	data.Delete_Mobile(id)
}
