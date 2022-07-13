package service

import (
	"DailyActivity/dto"
	"DailyActivity/entity"
	"DailyActivity/repository"
	"DailyActivity/service/activity"
	"github.com/mashingan/smapping"
	"log"
	"strconv"
)

type ActivityService interface {
	All(userID string) (*[]activity.ActivityResponse, error)
	CreateActivity(activityRequest dto.ActivityCreate, userID string) (*activity.ActivityResponse, error)
	IsCheckedIn(userId string) bool
	IsCheckedOut(userId string) bool
}

type activityService struct {
	activityRepository repository.ActivityRepository
}

func NewActivityService(activityRepository repository.ActivityRepository) ActivityService {
	return &activityService{
		activityRepository: activityRepository,
	}
}

func (c *activityService) All(userID string) (*[]activity.ActivityResponse, error) {
	acts, err := c.activityRepository.All(userID)
	if err != nil {
		return nil, err
	}

	acti := activity.NewActivityArrayResponse(acts)
	return &acti, nil
}

func (c *activityService) CreateActivity(activityRequest dto.ActivityCreate, userID string) (*activity.ActivityResponse, error) {
	act := entity.Activity{}
	err := smapping.FillStruct(&act, smapping.MapFields(&activityRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	act.UserID = uint64(id)
	p, err := c.activityRepository.InsertActivity(act)
	if err != nil {
		return nil, err
	}

	res := activity.NewActivityResponse(p)
	return &res, nil
}

func (c *activityService) IsCheckedIn(userID string) bool {
	res := c.activityRepository.IsCheckIn(userID)
	return res.Error == nil
}

func (c *activityService) IsCheckedOut(userID string) bool {
	res := c.activityRepository.IsCheckOut(userID)
	return res.Error == nil
}
