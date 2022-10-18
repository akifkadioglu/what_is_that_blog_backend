package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Email   string `json:"email"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
