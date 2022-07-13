package repository

import (
	"DailyActivity/entity"
	"gorm.io/gorm"
	"time"
)

type AttendanceRepository interface {
	CheckIn(attendance entity.Attendance) (entity.Attendance, error)
	CheckOut(attendance entity.Attendance) (entity.Attendance, error)
	IsDuplicate(userID string) (tx *gorm.DB)
	IsCheckedOut(userID string) (tx *gorm.DB)
}

type attendanceRepo struct {
	connection *gorm.DB
}

func NewAttendanceRepo(connection *gorm.DB) AttendanceRepository {
	return &attendanceRepo{
		connection: connection,
	}
}

func (db *attendanceRepo) All(userID string) ([]entity.Attendance, error) {
	attendances := []entity.Attendance{}
	db.connection.Preload("User").Where("user_id = ?", userID).Find(&attendances)
	return attendances, nil
}

func (db *attendanceRepo) IsDuplicate(userID string) (tx *gorm.DB) {
	var attendance entity.Attendance
	return db.connection.Where("user_id = ?", userID).
		Where("DATE(created_at) = ? AND attendance_type ='IN'", time.Now().Format("2006-01-02")).
		Take(&attendance)
}

func (db *attendanceRepo) IsCheckedOut(userID string) (tx *gorm.DB) {
	var attendance entity.Attendance
	return db.connection.Where("user_id = ?", userID).
		Where("DATE(created_at) = ? AND attendance_type ='OUT'", time.Now().Format("2006-01-02")).
		Take(&attendance)
}

func (db *attendanceRepo) CheckIn(attendance entity.Attendance) (entity.Attendance, error) {
	db.connection.Save(&attendance)
	db.connection.Preload("User").Find(&attendance)
	return attendance, nil
}

func (db *attendanceRepo) CheckOut(attendance entity.Attendance) (entity.Attendance, error) {
	db.connection.Save(&attendance)
	db.connection.Preload("User").Find(&attendance)
	return attendance, nil
}
