package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const SecretKey = "secret"

// Generate JWT
// Expects Issuer (string)
// Returns string, err
func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day JWT expire
	})

	return claims.SignedString([]byte(SecretKey))

}

// Parse JWT
// Expects cookie (string)
// Returns string, err
func ParseJWT(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}
