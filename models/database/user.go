package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Login        string `gorm:"type:varchar(20);unique"`
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
