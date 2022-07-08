package entity

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"time"
)

type attendanceType string

const (
	IN  attendanceType = "IN"
	OUT attendanceType = "OUT"
)

func (ct *attendanceType) Scan(value interface{}) error {
	*ct = attendanceType(value.([]byte))
	return nil
}

func (ct attendanceType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Attendance struct {
	ID             uint64         `gorm:"primary_key:auto_increment" json:"id"`
	UserID         uint64         `gorm:"not null" json:"-"`
	User           User           `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	AttendanceType attendanceType `gorm:"not null" json:"attendance_type"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
