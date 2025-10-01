package repository

import (
	"github.com/fauzan264/backend/dto/request"
	"github.com/fauzan264/backend/model"
	"gorm.io/gorm"
)

type feedRepository struct {
	db *gorm.DB
}

type FeedRepository interface {
	GetFeeds(requestSearch request.GetFeedsRequest) ([]model.Feed, error)
	CreateFeed(feed model.Feed) (model.Feed, error)
}

func NewFeedRepository (db *gorm.DB) *feedRepository {
	return &feedRepository{db}
}

func (r *feedRepository) GetFeeds(requestSearch request.GetFeedsRequest) ([]model.Feed, error) {
	var feeds []model.Feed

	if requestSearch.Page < 1 {
		requestSearch.Page = 1
	}
	if requestSearch.Limit < 1 {
		requestSearch.Limit = 10
	}

	offset := (requestSearch.Page - 1) * requestSearch.Limit

	query := r.db.Model(&model.Feed{})

	err := query.Limit(requestSearch.Limit).
					Offset(offset).
					Order("createdat DESC").
					Find(&feeds).Error

	if err != nil {
		return nil, err
	}

	return feeds, nil
}

func (r *feedRepository) CreateFeed(feed model.Feed) (model.Feed, error) {
	err := r.db.Create(&feed).Error
	if err != nil {
		return feed, err
	}

	return feed, nil
}
