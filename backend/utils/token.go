package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a JWT token with user information
func GenerateToken(id uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

// VerifyToken validates a JWT token and returns claims
func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// ExtractUserID extracts user ID from JWT claims
func ExtractUserID(claims *jwt.MapClaims) uint {
	if userID, ok := (*claims)["user_id"].(float64); ok {
		return uint(userID)
	}
	return 0
}

// ExtractUserRole extracts user role from JWT claims
func ExtractUserRole(claims *jwt.MapClaims) string {
	if role, ok := (*claims)["role"].(string); ok {
		return role
	}
	return ""
}
