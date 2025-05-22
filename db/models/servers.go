package models

import (
	"time"
)

type Server struct {
	ID        uint    `gorm:"primaryKey"`
	Hostname  string  `gorm:"not null;unique"`
	IPAddress string  `gorm:"not null"`
	Location  *string // opsiyonel alan, eklediysen
	CreatedAt time.Time
}
