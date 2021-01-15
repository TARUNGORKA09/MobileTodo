package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/TARUNGORKA09/MobileTodo/data"
	"github.com/gorilla/mux"
)

type MobileInfo struct {
	l *log.Logger
}

func NewMobile(m *log.Logger) *MobileInfo {
	return &MobileInfo{m}
}

func (p *MobileInfo) GetMobileInfo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")

	password := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if password[0] != "Basic" {
		fmt.Errorf("Not a valid Authorization", password)
	}

	pass, err := base64.StdEncoding.DecodeString(password[1])
	if err != nil {
		fmt.Errorf("unable to convert from base64 to string", password)
	}
	pass1 := strings.SplitN(string(pass), ":", 2)
	fmt.Fprint(rw, string(pass1[0]))
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	lp := data.GetMobile(id)
	err1 := lp.ToJSON(rw)
	//fmt.Fprintf(rw, " ", lp)
	if err1 != nil {
		http.Error(rw, "unable to marshall ", http.StatusInternalServerError)
	}

}
