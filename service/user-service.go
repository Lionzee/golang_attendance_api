package service

import (
	"DailyActivity/dto"
	"DailyActivity/entity"
	"DailyActivity/repository"
	"DailyActivity/service/user"
	"github.com/mashingan/smapping"
	"log"
)

//UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	FindUserByID(userID string) (*user.UserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) FindUserByID(userID string) (*user.UserResponse, error) {
	userData, err := service.userRepository.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	userResponse := user.UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&userData))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}
