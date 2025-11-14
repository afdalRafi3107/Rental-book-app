package entity

import "time"

type Book struct {
	ID          uint       `gorm:"primaryKey"`
	Title       string
	Author      string
	Description string
	CoverURL    string
	Price       int64
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
