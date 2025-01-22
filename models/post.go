package models

type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Category string `json:"category"`
}
