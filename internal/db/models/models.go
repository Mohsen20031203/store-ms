package models

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model `json:"_"`
	ID         int       `gorm:"primaryKey"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"type:text;not null"`
	CategoryID int       `gorm:"index"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

type Category struct {
	gorm.Model `json:"-"`
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(100);unique;not null"`
}
