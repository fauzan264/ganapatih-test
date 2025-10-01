package service

import (
	"time"

	"github.com/fauzan264/backend/dto/request"
	"github.com/fauzan264/backend/dto/response"
	"github.com/fauzan264/backend/model"
	repository "github.com/fauzan264/backend/repositories"
)

type feedService struct {
	feedRepository repository.FeedRepository
}

type FeedService interface {
	GetFeeds(requestSearch request.GetFeedsRequest) (response.FeedsResponse, error)
	CreateFeed(request request.CreateFeedRequest) (response.FeedResponse, error)
}

func NewFeedService(feedRepository repository.FeedRepository) FeedService {
	return &feedService{feedRepository}
}

func (s *feedService) GetFeeds(requestSearch request.GetFeedsRequest) (response.FeedsResponse, error) {
	feeds, err := s.feedRepository.GetFeeds(requestSearch)
	if err != nil {
		return response.FeedsResponse{}, err
	}

	feedResponses := make([]response.FeedResponse, 0)

	for _, feed := range feeds {
			feedResponses = append(feedResponses, response.FeedResponse{
					ID:        feed.ID,
					Userid:    feed.Userid,
					Content:   feed.Content,
					Createdat: feed.Createdat,
			})
	}

	response := response.FeedsResponse{
		Page: requestSearch.Page,
		Posts: feedResponses,
	}

	return response, nil
}

func (s *feedService) CreateFeed(request request.CreateFeedRequest) (response.FeedResponse, error) {
	feed := model.Feed{
		Userid:    request.Userid,
		Content:   request.Content,
		Createdat: time.Now(),
	}

	createdFeed, err := s.feedRepository.CreateFeed(feed)
	if err != nil {
		return response.FeedResponse{}, err
	}

	result := response.FeedResponse{
		ID:        createdFeed.ID,
		Userid:    createdFeed.Userid,
		Content:   createdFeed.Content,
		Createdat: createdFeed.Createdat,
	}

	return result, nil
}
