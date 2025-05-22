package models

import "time"

type WorkerConnection struct {
	ID          uint   `gorm:"primaryKey"`
	WorkerID    uint   `gorm:"not null"` // foreign key if needed later
	IP          string `gorm:"not null"`
	ConnectedAt time.Time
}
