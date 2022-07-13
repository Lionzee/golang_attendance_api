package service

import (
	"DailyActivity/entity"
	"DailyActivity/repository"
	"DailyActivity/service/attendance"
	"strconv"
)

type AttendanceService interface {
	CheckIn(userID string) (*attendance.AttendanceResponse, error)
	CheckOut(userID string) (*attendance.AttendanceResponse, error)
	IsDuplicate(userId string) bool
	IsCheckedIn(userId string) bool
	IsCheckedOut(userId string) bool
}

type attendanceService struct {
	attendanceRepo repository.AttendanceRepository
}

func NewAttendanceService(attendanceRepo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{
		attendanceRepo: attendanceRepo,
	}
}

func (service *attendanceService) IsDuplicate(userID string) bool {
	res := service.attendanceRepo.IsDuplicate(userID)
	return !(res.Error == nil)
}

func (service *attendanceService) IsCheckedIn(userID string) bool {
	res := service.attendanceRepo.IsDuplicate(userID)
	return res.Error == nil
}

func (service *attendanceService) IsCheckedOut(userID string) bool {
	res := service.attendanceRepo.IsCheckedOut(userID)
	return !(res.Error == nil)
}

func (service *attendanceService) CheckIn(userID string) (*attendance.AttendanceResponse, error) {
	att := entity.Attendance{}
	id, _ := strconv.ParseInt(userID, 0, 64)
	att.UserID = uint64(id)
	att.AttendanceType = "IN"
	a, err := service.attendanceRepo.CheckIn(att)
	if err != nil {
		return nil, err
	}

	res := attendance.NewCheckIn(a)
	return &res, nil
}

func (service *attendanceService) CheckOut(userID string) (*attendance.AttendanceResponse, error) {
	att := entity.Attendance{}
	id, _ := strconv.ParseInt(userID, 0, 64)
	att.UserID = uint64(id)
	att.AttendanceType = "OUT"
	a, err := service.attendanceRepo.CheckIn(att)
	if err != nil {
		return nil, err
	}

	res := attendance.NewCheckIn(a)
	return &res, nil
}
