package middleware

import (
	"auth-service/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware validates JWT tokens and extracts user info
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized - Missing or invalid token",
			})
		}

		tokenStr := strings.Split(authHeader, " ")[1]
		claims, err := jwt.VerifyToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized - Invalid token",
			})
		}

		// ✅ Store claims (user info) in Fiber context
		c.Locals("user", claims)

		return c.Next()
	}
}
