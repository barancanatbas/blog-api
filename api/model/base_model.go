package model

import "time"

type Base struct {
	ID        uint32    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
