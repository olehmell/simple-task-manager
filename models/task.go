package models

type Task struct {
	ID       uint16 `json:"id"`
	ColumnID uint16 `json:"column_id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
}
