package models

import (
	"time"
)

type Job struct {
	ID         uint   `gorm:"primaryKey"`
	WorkerID   uint   `gorm:"not null"`
	JobType    string `gorm:"not null"` // e.g. "SNMP", "Ping"
	Target     string `gorm:"not null"`
	Status     string `gorm:"default:queued"` // queued, running, done, failed
	CreatedAt  time.Time
	StartedAt  *time.Time
	FinishedAt *time.Time
}
