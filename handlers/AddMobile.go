package handlers

import (
	"fmt"
	"net/http"

	"github.com/TARUNGORKA09/MobileTodo/data"
)

type keyMobile struct {
}

func (m *MobileInfo) AddMobile(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("content-type", "application/json")

	mob := r.Context().Value(keyMobile{}).(*data.Mobile)
	m.l.Print("Data %#v", mob)
	data.Add_Mobile(mob)

}
