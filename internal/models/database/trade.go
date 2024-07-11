package database

import (
	"gorm.io/gorm"
	"time"
)

type Trade struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	BookName    string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:varchar(250)"`

	AuthorId uint
	User     User `gorm:"foreignKey:AuthorId;reference:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
