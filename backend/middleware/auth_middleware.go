package middleware

import (
	"strings"

	"github.com/fauzan264/backend/dto/request"
	service "github.com/fauzan264/backend/services"
	"github.com/fauzan264/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(authService service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeaderToken := c.Get("Authorization")
		if !strings.Contains(authHeaderToken, "Bearer") {
			return sendUnauthorizedResponse()
		}

		tokenString := ""
		arrayToken := strings.Split(authHeaderToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		jwtService := utils.NewJWTService()
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return sendUnauthorizedResponse()
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return sendUnauthorizedResponse()
		}

		userID := claims["id"].(float64)

		requestUser := request.GetUser{
			ID: int(userID),
		}

		user, err := authService.GetUserByID(requestUser.ID)
		if err != nil {
			return sendUnauthorizedResponse()
		}

		c.Locals("authUser", user)

		return c.Next()
	}
}

func sendUnauthorizedResponse() error {
	return fiber.NewError(fiber.StatusUnauthorized, "Please provide a valid authentication token")
}