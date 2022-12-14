package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Image    string `json:"image"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId int    `json:"author_id"`
	Author   User   `json:"author" gorm:"foreignKey:author_id; References:id"`
}
