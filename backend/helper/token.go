package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTTokenExpiry  = 7 * 24 * time.Hour
)


func GenerateAccessToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(JWTTokenExpiry).Unix(),
		"iat":     time.Now().Unix(),
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}


func VerifyAccessToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, ok := claims["type"].(string); !ok || tokenType != "access" {
			return nil, jwt.ErrTokenInvalidClaims
		}
		return &claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func ExtractUserID(claims *jwt.MapClaims) uint {
	if userID, ok := (*claims)["user_id"].(float64); ok {
		return uint(userID)
	}
	return 0
}

func ExtractUserRole(claims *jwt.MapClaims) string {
	if role, ok := (*claims)["role"].(string); ok {
		return role
	}
	return ""
}




func GenerateToken(id uint, role string) (string, error) {
	return GenerateAccessToken(id, role)
}

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	return VerifyAccessToken(tokenString)
}




