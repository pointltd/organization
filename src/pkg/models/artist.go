package models

type Artist struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Genre   string `json:"genre"`
	Country string `json:"country"`
}
