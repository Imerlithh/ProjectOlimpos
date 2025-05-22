package models

import (
	"gorm.io/gorm"
	"time"
)

type APIRequestLog struct {
	ID           uint    `gorm:"primaryKey"`
	Endpoint     string  `gorm:"not null"`
	Method       string  `gorm:"not null"`
	RequestBody  *string // nullable string
	ResponseCode int     `gorm:"not null"`
	ClientIP     string  `gorm:"not null"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"` // soft delete desteği opsiyonel
}
