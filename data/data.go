package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Mobile struct {
	General1          General          `json:"general"`
	Display_features1 Display_features `json:"display_features"`
	Os_Pro1           Os_Pro           `json:"os and processors"`
	Mem_Feat1         Mem_Feat         `json:"memory features"`
	Cam_feat1         Cam_feat         `json:"camera features"`
	Conn_Feat1        Conn_Feat        `json:"connection features"`
}

type General struct {
	Mod_num   int    `json:"Id"`
	Mod_name  string `json:"mod_name"`
	Color     string `json:"color"`
	Sim_type  string `json:"sim_type"`
	Touch_scr bool   `json:"touch_scr"`
}

type Display_features struct {
	Disp_size    float32 `json:"disp_size"`
	Resolution   int     `json:"resolution"`
	Display_type string  `json:"disp_type"`
}

type Os_Pro struct {
	Os       string `json:"os"`
	Pro_type string `json:"pro_type"`
	Pro_core string `json:"pro_core"`
}

type Mem_Feat struct {
	Int_stor    int   `json:"int_stor"`
	Ram         int   `json:"ram"`
	Expand_stor int32 `json:"expand_stor"`
}

type Cam_feat struct {
	Prim_cam_avail   bool   `json:"prim_cam_avail"`
	Prim_cam         string `json:"prim_cam"`
	Second_cam_avail bool   `json:"second_cam_avail"`
	Second_cam       int    `json:"seond_cam"`
}

type Conn_Feat struct {
	Net_type  string  `json:"net_type"`
	_3G       bool    `json:"3G"`
	GPRS      bool    `json:"GPRS"`
	Bluetooth bool    `json:"bluetooth"`
	Blu_ver   float32 `json:"blu_ver"`
	Wifi      bool    `json:"wifi"`
	GPS       bool    `json:"gps"`
}

func GetMobile(id int) *Mobile {

	_, pos, err := findMobile(id)
	if err != nil {
		fmt.Print("Unable to find mobile")
	}
	return mobileList[pos]

}

func Add_Mobile(m *Mobile) {

	mob := mobileList[len(mobileList)-1]
	m.General1.Mod_num = mob.General1.Mod_num + 1
	mobileList = append(mobileList, m)

}

func findMobile(id int) (*Mobile, int, error) {
	for i, p := range mobileList {
		if p.General1.Mod_num == id {
			return p, i, nil
		}
	}
	return nil, -1, http.ErrBodyNotAllowed
}

func (m *Mobile) ToJSON(w io.Writer) error {
	s := json.NewEncoder(w)
	return s.Encode(m)
}

func (m *Mobile) ToDATA(r io.Reader) error {
	s := json.NewDecoder(r)
	return s.Decode(m)
}

var mobileList = []*Mobile{
	{
		General1: General{
			Mod_num:   1,
			Mod_name:  "Xiaomi",
			Color:     "white",
			Sim_type:  "Dual",
			Touch_scr: true,
		},
		Display_features1: Display_features{
			Disp_size:    6.5,
			Display_type: "AMOLED",
			Resolution:   4500,
		},
		Os_Pro1: Os_Pro{
			Os:       "Android",
			Pro_core: "octa core",
			Pro_type: "MediaTek",
		},
		Mem_Feat1: Mem_Feat{
			Int_stor:    64,
			Expand_stor: 512,
			Ram:         8,
		},
		Cam_feat1: Cam_feat{
			Prim_cam_avail:   true,
			Prim_cam:         "48 x 2 x 5",
			Second_cam_avail: true,
			Second_cam:       48,
		},
		Conn_Feat1: Conn_Feat{
			Net_type:  "4G",
			_3G:       true,
			GPRS:      true,
			Bluetooth: true,
			Blu_ver:   5.1,
			Wifi:      true,
			GPS:       true,
		},
	},
}
