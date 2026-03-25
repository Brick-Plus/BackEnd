package security

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ClaimsAdmin struct {
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

type Claims struct {
	UserId string `json:"id"`
	jwt.RegisteredClaims
}

func VerifyAdminToken(tokenString string) (*ClaimsAdmin, error) {
	godotenv.Load()
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsAdmin{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimsAdmin)
	if !ok || !token.Valid || !claims.IsAdmin {
		return nil, errors.New("invalid token, not an admin")
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