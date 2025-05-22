package models

import "time"

type WorkerJoinRequest struct {
	ID        uint   `gorm:"primaryKey"`
	Hostname  string `gorm:"not null"`
	IP        string `gorm:"not null"`
	Token     string `gorm:"not null;unique"`
	Status    string `gorm:"default:pending"` // pending, approved, rejected
	CreatedAt time.Time
}
