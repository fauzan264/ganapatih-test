package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fauzan264/backend/dto/request"
	"github.com/fauzan264/backend/dto/response"
	service "github.com/fauzan264/backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService}
}

func (h *authHandler) RegisterUser(c *fiber.Ctx) error {
	var request request.RegisterRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Register user failed")
	}

	registerUserResponse, err := h.authService.RegisterUser(request)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
				case "23505":
					return fiber.NewError(http.StatusConflict, fmt.Sprintf("Error: username \"%s\" already taken.", request.Username))
				}
		}
		
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to register user")
	}

	return c.Status(fiber.StatusOK).JSON(registerUserResponse)
}

func (h *authHandler) LoginUser(c *fiber.Ctx) error {
	var request request.LoginRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Login user failed")
	}

	userResponse, err := h.authService.LoginUser(request)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid username or password")
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Token: userResponse.Token,
	})
}

func (h *authHandler) SessionUser(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	if authUser == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	user, ok := authUser.(response.UserResponse)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
	}

	userResponse, _ := h.authService.SessionUser(user.ID)

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Data: userResponse,
	})
}
