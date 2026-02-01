package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

const (
	AccessTokenCookie  = "access_token"
	RefreshTokenCookie = "refresh_token"
)

func GenerateAccessToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(AccessTokenExpiry).Unix(),
		"iat":     time.Now().Unix(),
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
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

type CookieConfig struct {
	Secure   bool
	Domain   string
	SameSite string
}

func GetCookieConfig() CookieConfig {
	isProduction := os.Getenv("GO_ENV") == "production"

	return CookieConfig{
		Secure:   isProduction,
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		SameSite: "Lax",
	}
}

func SetAccessTokenCookie(c *fiber.Ctx, token string) {
	config := GetCookieConfig()

	c.Cookie(&fiber.Cookie{
		Name:     AccessTokenCookie,
		Value:    token,
		Expires:  time.Now().Add(AccessTokenExpiry),
		HTTPOnly: true,
		Secure:   config.Secure,
		SameSite: config.SameSite,
		Path:     "/",
	})
}

func SetRefreshTokenCookie(c *fiber.Ctx, token string) {
	config := GetCookieConfig()

	c.Cookie(&fiber.Cookie{
		Name:     RefreshTokenCookie,
		Value:    token,
		Expires:  time.Now().Add(RefreshTokenExpiry),
		HTTPOnly: true,
		Secure:   config.Secure,
		SameSite: config.SameSite,
		Path:     "/api/v1/auth",
	})
}

func ClearAuthCookies(c *fiber.Ctx) {
	config := GetCookieConfig()

	c.Cookie(&fiber.Cookie{
		Name:     AccessTokenCookie,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   config.Secure,
		SameSite: config.SameSite,
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     RefreshTokenCookie,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   config.Secure,
		SameSite: config.SameSite,
		Path:     "/api/v1/auth",
	})
}

func GetAccessTokenFromCookie(c *fiber.Ctx) string {
	return c.Cookies(AccessTokenCookie)
}

func GetRefreshTokenFromCookie(c *fiber.Ctx) string {
	return c.Cookies(RefreshTokenCookie)
}

func GenerateToken(id uint, role string) (string, error) {
	return GenerateAccessToken(id, role)
}

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	return VerifyAccessToken(tokenString)
}
