package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type mobile struct {
	general          General
	display_features Display_features
	os_Pro           Os_Pro
	mem_Feat         Mem_Feat
	cam_feat         Cam_feat
	conn_Feat        Conn_Feat
}

type General struct {
	Mod_num   int `json:"Id"`
	Mod_name  string
	Color     string
	Sim_type  string
	Touch_scr bool
}

type Display_features struct {
	Disp_size    float32
	Resolution   int
	Display_type string
}

type Os_Pro struct {
	Os       string
	Pro_type string
	Pro_core string
}

type Mem_Feat struct {
	Int_stor    int
	Ram         int
	Expand_stor int32
}

type Cam_feat struct {
	Prim_cam_avail   bool
	Prim_cam         string
	Second_cam_avail bool
	Second_cam       int
}

type Conn_Feat struct {
	Net_type  string
	_3G       bool
	GPRS      bool
	Bluetooth bool
	Blu_ver   float32
	wifi      bool
	GPS       bool
}

var mobileList = []*mobile{
	{
		general: General{
			Mod_num:   1,
			Mod_name:  "Xiaomi",
			Color:     "white",
			Sim_type:  "Dual",
			Touch_scr: true,
		},
		display_features: Display_features{
			Disp_size:    6.5,
			Display_type: "AMOLED",
			Resolution:   4500,
		},
		os_Pro: Os_Pro{
			Os:       "Android",
			Pro_core: "octa core",
			Pro_type: "MediaTek",
		},
		mem_Feat: Mem_Feat{
			Int_stor:    64,
			Expand_stor: 512,
			Ram:         8,
		},
		cam_feat: Cam_feat{
			Prim_cam_avail:   true,
			Prim_cam:         "48 x 2 x 5",
			Second_cam_avail: true,
			Second_cam:       48,
		},
		conn_Feat: Conn_Feat{
			Net_type:  "4G",
			_3G:       true,
			GPRS:      true,
			Bluetooth: true,
			Blu_ver:   5.1,
			wifi:      true,
			GPS:       true,
		},
	},
}

func (m mobile) getMobile(id int) *mobile {

	_, pos, err := findProduct(id)
	if err != nil {
		fmt.Print("Unable to find mobile")
	}
	return mobileList[pos]

}

func findProduct(id int) (*mobile, int, error) {
	for i, p := range mobileList {
		if p.general.Mod_num == id {
			return p, i, nil
		}
	}
	return nil, -1, http.ErrBodyNotAllowed
}

func (m *mobile) ToJSON(w io.Writer) error {
	s := json.NewEncoder(w)
	return s.Encode(m)
}

func (m *mobile) ToDATA(r io.Reader) error {
	s := json.NewDecoder(r)
	return s.Decode(m)
}
