package database

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
