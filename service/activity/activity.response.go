package activity

import (
	"DailyActivity/entity"
	"DailyActivity/service/user"
	"time"
)

type ActivityResponse struct {
	ID          uint64            `json:"id"`
	User        user.UserResponse `json:"user,omitempty"`
	Description string            `json:"description"`
	CreatedAt   time.Time         `json:"created_at"`
}

func NewActivityLog(activity entity.Activity) ActivityResponse {
	return ActivityResponse{
		ID:        activity.ID,
		CreatedAt: activity.CreatedAt,
		User:      user.NewUserResponse(activity.User),
	}
}
