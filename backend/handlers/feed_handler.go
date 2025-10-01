package handlers

import (
	"github.com/fauzan264/backend/dto/request"
	"github.com/fauzan264/backend/dto/response"
	service "github.com/fauzan264/backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


type feedHandler struct {
	feedService service.FeedService
}

func NewFeedHandler(feedService service.FeedService) *feedHandler {
	return &feedHandler{feedService}
}

func (h *feedHandler) CreateFeed(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	if authUser == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	var req request.CreateFeedRequest
	err := c.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	req.Userid = user.ID
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Content" && err.Tag() == "max" {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": "Content exceeds maximum length of 200 characters",
					"code":  422,
				})
			}
		}
		return fiber.NewError(fiber.StatusBadRequest, "Create feed failed")
	}

	feedResponse, err := h.feedService.CreateFeed(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create feed")
	}

	return c.Status(fiber.StatusCreated).JSON(response.Response{
		Data: feedResponse,
	})
}

func (h *feedHandler) GetFeeds(c *fiber.Ctx) error {
	var req request.GetFeedsRequest
	err := c.QueryParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Get feeds failed")
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 {
		req.Limit = 10
	}

	feedsResponse, err := h.feedService.GetFeeds(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get feeds")
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Data: feedsResponse,
	})
}