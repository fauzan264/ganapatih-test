package request

type GetFeedsRequest struct {
	Page  int `json:"page" validate:"omitempty,min=1"`
	Limit int `json:"limit" validate:"omitempty,min=1,max=100"`
}

type CreateFeedRequest struct {
	Userid 		int 		`json:"userid" validate:"required"`
	Content 	string 	`json:"content"  validate:"required,min=1,max=200"`
}