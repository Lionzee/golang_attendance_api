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

func NewActivityResponse(activity entity.Activity) ActivityResponse {
	return ActivityResponse{
		ID:          activity.ID,
		Description: activity.Description,
		CreatedAt:   activity.CreatedAt,
		User:        user.NewUserResponse(activity.User),
	}
}

func NewActivityArrayResponse(activities []entity.Activity) []ActivityResponse {
	activityRes := []ActivityResponse{}
	for _, activity := range activities {
		p := ActivityResponse{
			ID:          activity.ID,
			Description: activity.Description,
			CreatedAt:   activity.CreatedAt,
			User:        user.NewUserResponse(activity.User),
		}
		activityRes = append(activityRes, p)
	}
	return activityRes
}
