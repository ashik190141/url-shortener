package model

import (
	"time"
)

type URL struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	LongURL   string    `gorm:"type:text;not null"`
	ShortURL  string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	UserID    int       `gorm:"index;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}