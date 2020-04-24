package usercase

import (
	"github.com/hoangvantuan/go-base/domain/model"
	"github.com/hoangvantuan/go-base/domain/repo"
	"github.com/hoangvantuan/go-base/domain/service"
	"github.com/hoangvantuan/go-base/helper"
	"github.com/pkg/errors"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserUsecase interface {
	FindAll() ([]*User, error)
}

type userUsecase struct {
	userRepo    repo.UserRepository
	userService *service.UserService
}

func NewUserUsercase(userRepo repo.UserRepository, userService *service.UserService) UserUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		userService: userService,
	}
}

func (u *userUsecase) FindAll() ([]*User, error) {
	us, err := u.userRepo.FindAll()
	if err != nil {
		return nil, helper.Pe(errors.Wrapf(err, "error while while find all user"))
	}

	return toUser(us), nil
}

func toUser(us []*model.User) []*User {
	var users []*User
	for _, u := range us {
		users = append(users, &User{
			ID:   u.ID,
			Name: u.Name,
		})
	}

	return users
}
