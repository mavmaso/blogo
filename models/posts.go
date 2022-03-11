package models

type Posts struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
