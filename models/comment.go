package models

type Comment struct {
	ID     uint16 `json:"id"`
	TaskID uint16 `json:"task_id"`
	Text   string `json:"text"`
}
