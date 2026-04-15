package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims structure unique pour tous les utilisateurs
type Claims struct {
	UserId  string `json:"id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (string, error) {
	return generateTokenWithRole(userId, false)
}

func GenerateAdminToken(userId string) (string, error) {
	return generateTokenWithRole(userId, true)
}

func generateTokenWithRole(userId string, isAdmin bool) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	claims := &Claims{
		UserId:  userId,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}
	return signed, nil
}

func VerifyAdminToken(tokenString string) (*Claims, error) {
	claims, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	if !claims.IsAdmin {
		return nil, errors.New("User isn't an admin !")
	}

	return claims, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}