package security

import (
	model "GDN-delivery-management/db/sql"
	"github.com/dgrijalva/jwt-go"
	"time"

	"fmt"
)

const SecretKey = "fshjofjsdfo8oi3wyuf98wyu9876uhzxiou#@%%"

func GenToken(user model.User) (string, *JwtCustomClaims, error) {
	claims := &JwtCustomClaims{
		UserId: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", nil, err
	}
	return result, claims, nil
}

func GenRefreshtoken(user model.User) (string, *JwtCustomClaims, error) {
	claims := &JwtCustomClaims{
		UserId: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resut, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", nil, err
	}
	return resut, claims, nil
}

func VerifyToken(accessToken string) (*JwtCustomClaims, error) {
	// Parse the access token string
	token, err := jwt.ParseWithClaims(accessToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to sign the token
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

type JwtCustomClaims struct {
	UserId string
	Email  string
	jwt.StandardClaims
}
