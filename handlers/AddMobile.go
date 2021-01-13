package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TARUNGORKA09/MobileTodo/data"
)

type keyMobile struct {
}

func (m *MobileInfo) AddMobile(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("content-type", "application/json")
	mob := json.NewDecoder(r.Body)
	var t data.Mobile
	mob.Decode(&t)
	data.Add_Mobile(&t)

}
