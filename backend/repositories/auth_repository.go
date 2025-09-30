package repository

import (
	"github.com/fauzan264/backend/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	Register(user model.User) (model.User, error)
	Login(user model.User) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserByID(ID int) (model.User, error)
	Session(user model.User) (model.User, error)
}

func NewAuthRepository (db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *authRepository) Login(user model.User) (model.User, error) {
	err := r.db.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *authRepository) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *authRepository) GetUserByID(ID int) (model.User, error) {
	user := model.User{}
	err := r.db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *authRepository) Session(user model.User) (model.User, error) {
	err := r.db.First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
