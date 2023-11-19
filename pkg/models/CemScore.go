package models

type CemScore struct {
	Timescale string `json:"timescale"`
	Osat int `json:"osat"`
	Taste int `json:"taste"`
	Speed int `json:"speed"`
	Ace int `json:"ace"`
	Cleanliness int `json:"cleanliness"`
	Accuracy int `json:"accuracy"`
}