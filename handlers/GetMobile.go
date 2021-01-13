package handlers

import (
	"log"
	"net/http"
	"strconv"

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
