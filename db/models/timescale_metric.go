package models

import "time"

type TimescaleMetric struct {
	ID        uint      `gorm:"primaryKey"`
	HostID    uint      `gorm:"not null"`
	Metric    string    `gorm:"not null"` // cpu_usage, ram_usage, etc.
	Value     float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"index"`
}
