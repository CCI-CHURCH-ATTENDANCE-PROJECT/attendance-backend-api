package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"cci-api/internal/dto"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash checks if a password matches its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateUserID generates a unique user ID with CCIMRB prefix
func GenerateUserID() (string, error) {
	// Generate a random number
	max := big.NewInt(99999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// Format with leading zeros to ensure 5 digits
	return fmt.Sprintf("CCIMRB-%05d", n.Int64()), nil
}

// GenerateRandomToken generates a random token
func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// JWTClaims represents JWT claims
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token
func GenerateJWT(userID, email string, admin bool, secret string, expiry time.Duration) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Admin:  admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cci-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateJWT validates a JWT token and returns claims
func ValidateJWT(tokenString, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// StringToInt converts string to int with default value
func StringToInt(s string, defaultValue int) int {
	if s == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return value
}

// CalculatePagination calculates pagination values
func CalculatePagination(page, limit, total int) (int, int, int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	totalPages := (total + limit - 1) / limit
	offset := (page - 1) * limit

	return page, limit, offset, totalPages
}

// NewPagination creates a new pagination object
func NewPagination(page, limit, total int) dto.Pagination {
	_, _, _, totalPages := CalculatePagination(page, limit, total)
	return dto.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}

// GetLagosTime returns current time in Lagos timezone (UTC+1)
func GetLagosTime() time.Time {
	// Lagos is UTC+1
	location, _ := time.LoadLocation("Africa/Lagos")
	return time.Now().In(location)
}

// IsValidPassword checks if password meets requirements
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false
	specialChars := "!@#$%^&*()-_=+[]{}|;:',.<>?/~`"

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case strings.ContainsRune(specialChars, char):
			hasSpecial = true
		}

	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
