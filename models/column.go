package models

type Column struct {
	ID     uint16 `json:"id"`
	DeskID uint16 `json:"desk_id"`
	Name   string `json:"name"`
}
