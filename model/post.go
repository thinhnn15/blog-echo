package model

import "time"

type Post struct {
	Id int `json:"-" db:"id, omitempty"`
	Title string `json:"title" db:"title, omitempty"`
	Content string `json:"content" db:"content, omitempty"`
	CreateAt time.Time `json:"-" db:"created_at, omitempty"`
}
