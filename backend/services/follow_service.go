package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/fauzan264/backend/dto/response"
	"github.com/fauzan264/backend/model"
	repository "github.com/fauzan264/backend/repositories"
)

type followService struct {
	followRepository repository.FollowRepository
	userRepository repository.UserRepository
}

type FollowService interface {
	FollowUser(followerID, followedID int) (response.FollowResponse, error)
	UnfollowUser(followerID, followedID int) (response.FollowResponse, error)
}

func NewFollowService(followRepository repository.FollowRepository, userRepository repository.UserRepository) FollowService {
	return &followService{followRepository, userRepository}
}

func (s *followService) FollowUser(followerID, followedID int) (response.FollowResponse, error) {
	if followerID == followedID {
		return response.FollowResponse{}, errors.New("you cannot follow yourself")
	}

	userExists, err := s.userRepository.UserExists(followedID)
	if err != nil {
		return response.FollowResponse{}, err
	}

	if !userExists {
		return response.FollowResponse{}, errors.New("user not found")
	}

	isFollowing, err := s.followRepository.IsFollowing(followerID, followedID)
	if err != nil {
		return response.FollowResponse{}, err
	}

	if isFollowing {
		return response.FollowResponse{}, errors.New("you are already following this user")
	}

	follow := model.Follow{
		FollowerID: followerID,
		FollowedID: followedID,
		CreatedAt:  time.Now(),
	}

	err = s.followRepository.FollowUser(follow)
	if err != nil {
		return response.FollowResponse{}, err
	}

	response := response.FollowResponse{
		Message: fmt.Sprintf("you are now following user %d", followedID),
	}

	return response, nil
}

func (s *followService) UnfollowUser(followerID, followedID int) (response.FollowResponse, error) {
	userExists, err := s.userRepository.UserExists(followedID)
	if err != nil {
		return response.FollowResponse{}, err
	}

	if !userExists {
		return response.FollowResponse{}, errors.New("user not found")
	}
	
	isFollowing, err := s.followRepository.IsFollowing(followerID, followedID)
	if err != nil {
		return response.FollowResponse{}, err
	}

	if !isFollowing {
		return response.FollowResponse{}, errors.New("you are not following this user")
	}

	err = s.followRepository.UnfollowUser(followerID, followedID)
	if err != nil {
		return response.FollowResponse{}, err
	}

	response := response.FollowResponse{
		Message: fmt.Sprintf("you unfollowed user %d", followedID),
	}

	return response, nil
}