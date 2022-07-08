package entity

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	UserID      uint64    `gorm:"not null" json:"-"`
	User        User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
