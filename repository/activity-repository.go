package repository

import (
	"DailyActivity/entity"
	"gorm.io/gorm"
	"time"
)

type ActivityRepository interface {
	All(userID string) ([]entity.Activity, error)
	InsertActivity(product entity.Activity) (entity.Activity, error)
	IsCheckIn(userID string) (tx *gorm.DB)
	IsCheckOut(userID string) (tx *gorm.DB)
}

type activityRepository struct {
	connection *gorm.DB
}

func NewActivityRepository(connection *gorm.DB) ActivityRepository {
	return &activityRepository{
		connection: connection,
	}
}

func (c *activityRepository) All(userID string) ([]entity.Activity, error) {
	activities := []entity.Activity{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&activities)
	return activities, nil
}

func (c *activityRepository) IsCheckIn(userID string) (tx *gorm.DB) {
	var attendance entity.Attendance
	return c.connection.Where("user_id = ?", userID).
		Where("DATE(created_at) = ? AND attendance_type ='IN'", time.Now().Format("2006-01-02")).
		Take(&attendance)
}

func (c *activityRepository) IsCheckOut(userID string) (tx *gorm.DB) {
	var attendance entity.Attendance
	return c.connection.Where("user_id = ?", userID).
		Where("DATE(created_at) = ? AND attendance_type ='OUT'", time.Now().Format("2006-01-02")).
		Take(&attendance)
}

func (c *activityRepository) InsertActivity(activity entity.Activity) (entity.Activity, error) {
	c.connection.Save(&activity)
	c.connection.Preload("User").Find(&activity)
	return activity, nil
}
