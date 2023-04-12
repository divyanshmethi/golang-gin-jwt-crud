package models

import (
	"RESTAPI_Gin/db"
	"gorm.io/gorm"
)

// Entry struct contains gorm.Model struct along with Content and userID fiels
type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := db.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, nil
	}
	return entry, nil
}
