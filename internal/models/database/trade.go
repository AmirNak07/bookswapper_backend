package database

import (
	"gorm.io/gorm"
	"time"
)

type Trade struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	BookId      uint
	Book        Book `gorm:"foreignKey:BookId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
