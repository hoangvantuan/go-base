package repo

import (
	"github.com/hoangvantuan/go-base/domain/model"
	"github.com/hoangvantuan/go-base/helper"
	"github.com/hoangvantuan/go-base/logger"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// UserRepository is interface interactive with other layer
type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByName(name string) (*model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// Migrate
	if err := model.MigrateUser(db); err != nil {
		logger.Error("failed to migrate data for table user %s", err)
	}

	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) FindAll() ([]*model.User, error) {
	users := []*model.User{}
	u.DB.Find(&users)
	if u.DB.Error != nil {
		return nil, helper.Pe(errors.Wrapf(u.DB.Error, "error while find all user"))
	}

	if u.DB.RecordNotFound() {
		return nil, nil
	}

	return users, nil
}

func (u *userRepository) FindByName(name string) (*model.User, error) {
	user := &model.User{}
	u.DB.Find(user)
	if u.DB.Error != nil {
		return nil, helper.Pe(errors.Wrapf(u.DB.Error, "error while find user %s by name", name))
	}

	if u.DB.RecordNotFound() {
		return nil, nil
	}

	return user, nil
}
