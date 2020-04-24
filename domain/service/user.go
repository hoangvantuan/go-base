package service

import (
	"github.com/hoangvantuan/go-base/domain/repo"
	"github.com/hoangvantuan/go-base/helper"
)

type UserService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Dupplicated(name string) error {
	user, err := s.userRepo.FindAll()
	if user != nil {
		return helper.Pef("%s already exits", name)
	}
	if err != nil {
		return helper.Pe(err)
	}

	return nil
}
