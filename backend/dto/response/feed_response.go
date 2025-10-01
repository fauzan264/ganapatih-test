package response

import "time"

type FeedResponse struct {
	ID					int				`json:"id"`
	Userid			int				`json:"user_id"`
	Content			string		`json:"content"`
	Createdat		time.Time	`json:"createdat"`
}

type FeedsResponse struct {
	Page				int 						`json:"page"`
	Posts				[]FeedResponse	`json:"posts"`
}