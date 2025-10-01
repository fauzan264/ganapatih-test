package handlers

import (
	"strconv"

	"github.com/fauzan264/backend/dto/response"
	service "github.com/fauzan264/backend/services"
	"github.com/gofiber/fiber/v2"
)

type followHandler struct {
	followService service.FollowService
}

func NewFollowHandler(followService service.FollowService) *followHandler {
	return &followHandler{
		followService: followService,
	}
}

func (h *followHandler) FollowUser(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	if authUser == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	userIDParam := c.Params("userid")
	followedID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	followResponse, err := h.followService.FollowUser(user.ID, followedID)
	if err != nil {
		switch err.Error() {
		case "you cannot follow yourself":
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		case "user not found":
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		case "you are already following this user":
			return fiber.NewError(fiber.StatusConflict, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to follow user")
		}
	}

	return c.Status(fiber.StatusOK).JSON(followResponse)
}

func (h *followHandler) UnfollowUser(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	if authUser == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	userIDParam := c.Params("userid")
	followedID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	unfollowResponse, err := h.followService.UnfollowUser(user.ID, followedID)
	if err != nil {
		switch err.Error() {
		case "user not found":
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		case "you are not following this user":
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to unfollow user")
		}
	}

	return c.Status(fiber.StatusOK).JSON(unfollowResponse)
}