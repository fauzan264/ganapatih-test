package repository

import (
	"github.com/fauzan264/backend/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Register(user model.User) (model.User, error)
	Login(user model.User) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserByID(ID int) (model.User, error)
	Session(user model.User) (model.User, error)
	UserExists(userID int) (bool, error)
}

func NewUserRepository (db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Login(user model.User) (model.User, error) {
	err := r.db.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(ID int) (model.User, error) {
	user := model.User{}
	err := r.db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Session(user model.User) (model.User, error) {
	err := r.db.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UserExists(userID int) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("id = ?", userID).Count(&count).Error
	return count > 0, err
}