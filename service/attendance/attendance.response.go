package attendance

import (
	"DailyActivity/entity"
	"DailyActivity/service/user"
	"time"
)

type AttendanceResponse struct {
	ID        uint64            `json:"id"`
	Type      string            `json:"attendance_type"`
	CreatedAt time.Time         `json:"created_at"`
	User      user.UserResponse `json:"user,omitempty"`
}

func NewCheckIn(attendance entity.Attendance) AttendanceResponse {
	return AttendanceResponse{
		ID:        attendance.ID,
		CreatedAt: attendance.CreatedAt,
		Type:      string(attendance.AttendanceType),
		User:      user.NewUserResponse(attendance.User),
	}
}
