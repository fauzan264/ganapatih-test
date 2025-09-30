package service

import (
	"errors"

	"github.com/fauzan264/backend/dto/request"
	"github.com/fauzan264/backend/dto/response"
	"github.com/fauzan264/backend/model"
	repository "github.com/fauzan264/backend/repositories"
	"github.com/fauzan264/backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepository repository.AuthRepository
}

type AuthService interface {
	RegisterUser(request request.RegisterRequest) (response.RegisterResponse, error)
	LoginUser(request request.LoginRequest) (response.LoginResponse, error)
	SessionUser(ID int) (response.SessionResponse, error)
	GetUserByID(ID int) (response.UserResponse, error)
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository}
}

func (s *authService) RegisterUser(request request.RegisterRequest) (response.RegisterResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.RegisterResponse{}, errors.New("failed to hash password")
	}

	userData := model.User{
		Username: request.Username,
		PasswordHash: string(passwordHash),
	}

	user, err := s.authRepository.Register(userData)
	if err != nil {
		return response.RegisterResponse{}, err
	}

	userResponse := response.RegisterResponse{
		ID: user.ID,
		Username: user.Username,
	}

	return userResponse, nil
}

func (s *authService) LoginUser(request request.LoginRequest) (response.LoginResponse, error) {
	user, err := s.authRepository.GetUserByUsername(request.Username)
	
	if err != nil {
		return response.LoginResponse{}, errors.New("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return response.LoginResponse{}, errors.New("Invalid email or password")
	}

	token, err := utils.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		return response.LoginResponse{}, err
	}

	response := response.LoginResponse{Token: token}
	
	return response, nil
}

func (s *authService) SessionUser(ID int) (response.SessionResponse, error) {
	user, err := s.authRepository.GetUserByID(ID)
	if err != nil {
		return response.SessionResponse{}, errors.New("User Not Found")
	}
	
	response := response.SessionResponse{
		ID: user.ID,
		Username: user.Username,
	}

	return response, nil
}

func (s *authService) GetUserByID(ID int) (response.UserResponse, error) {
	user, err := s.authRepository.GetUserByID(ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	response := response.UserResponse{
		ID: user.ID,
		Username: user.Username,
	}

	return response, nil
}