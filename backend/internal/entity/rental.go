package entity

import "time"

type Rental struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	BookID    uint
	Status    string // pending, paid, returned
	Amount    int64
	CreatedAt time.Time
}
