package repository

import (
	"github.com/fauzan264/backend/model"
	"gorm.io/gorm"
)

type followRepository struct {
	db *gorm.DB
}

type FollowRepository interface {
	FollowUser(follow model.Follow) error
	UnfollowUser(followerID, followedID int) error
	IsFollowing(followerID, followedID int) (bool, error)
}

func NewFollowRepository(db *gorm.DB) *followRepository {
	return &followRepository{db: db}
}

func (r *followRepository) FollowUser(follow model.Follow) error {
	return r.db.Create(&follow).Error
}

func (r *followRepository) UnfollowUser(followerID, followedID int) error {
	return r.db.Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Delete(&model.Follow{}).Error
}

func (r *followRepository) IsFollowing(followerID, followedID int) (bool, error) {
	var count int64
	err := r.db.Model(&model.Follow{}).
		Where("follower_id = ? AND followed_id = ?", followerID, followedID).
		Count(&count).Error
	
	return count > 0, err
}