package models

import "time"

type Article struct {
	ID      uint   `json:"id" gorm:"primary_key"` // kalo primary key -> auto increment
	Title   string `json:"title"`
	Content string `json:"content"`

	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}
