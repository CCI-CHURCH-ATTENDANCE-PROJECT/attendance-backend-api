package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"cci-api/internal/config"
	"cci-api/internal/dto"
	"cci-api/internal/utils"

	"github.com/labstack/echo/v4"
)

// Helper to get user_id safely from context
func GetUserID(c echo.Context) (string, bool) {
	userIDValue := c.Get("user_id")
	userID, ok := userIDValue.(string)
	if !ok || userID == "" {
		return "", false
	}
	return userID, true
}

// JWTMiddleware validates JWT tokens
func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, dto.APIResponse{
					Success: false,
					Error: &dto.ErrorInfo{
						Code:    "MISSING_TOKEN",
						Message: "Authorization token is required for authentication",
					},
				})
			}

			// Check if it starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, dto.APIResponse{
					Success: false,
					Error: &dto.ErrorInfo{
						Code:    "INVALID_TOKEN_FORMAT",
						Message: "Authorization token must be in Bearer format",
					},
				})
			}

			// Extract token
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate token
			claims, err := utils.ValidateJWT(token, cfg.JWTSecret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, dto.APIResponse{
					Success: false,
					Error: &dto.ErrorInfo{
						Code:    "INVALID_TOKEN",
						Message: "Invalid or expired token",
					},
				})
			}

			// Store user info in context
			fmt.Println("JWT middleware claims.UserID:", claims.UserID)
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("admin", claims.Admin)

			// Always check for user_id before proceeding
			if _, ok := GetUserID(c); !ok {
				return c.JSON(http.StatusUnauthorized, dto.APIResponse{
					Success: false,
					Error: &dto.ErrorInfo{
						Code:    "INVALID_USER_ID",
						Message: "User ID missing or invalid in token.",
					},
				})
			}

			return next(c)
		}
	}
}

// AdminMiddleware checks if user is admin
func AdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			admin, ok := c.Get("admin").(bool)
			if !ok || !admin {
				return c.JSON(http.StatusForbidden, dto.APIResponse{
					Success: false,
					Error: &dto.ErrorInfo{
						Code:    "INSUFFICIENT_PRIVILEGES",
						Message: "You do not seem to have the Admin privileges required to acess this resource.",
					},
				})
			}
			return next(c)
		}
	}
}

// CORSMiddleware handles CORS
func CORSMiddleware(origins string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get("Origin")

			// Check if origin is allowed
			allowedOrigins := strings.Split(origins, ",")
			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if strings.TrimSpace(allowedOrigin) == origin {
					allowed = true
					break
				}
			}

			if allowed {
				c.Response().Header().Set("Access-Control-Allow-Origin", origin)
			}

			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusOK)
			}

			return next(c)
		}
	}
}

// SecurityHeadersMiddleware adds security headers
func SecurityHeadersMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("X-Content-Type-Options", "nosniff")
			c.Response().Header().Set("X-Frame-Options", "DENY")
			c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
			c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			c.Response().Header().Set("Content-Security-Policy", "default-src 'self'")
			return next(c)
		}
	}
}
