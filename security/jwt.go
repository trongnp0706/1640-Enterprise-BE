package security

import (
	"time"
	model "GDN-delivery-management/db/sql"
	"github.com/dgrijalva/jwt-go"

	"fmt"
)

const SecretKey = "fshjofjsdfo8oi3wyuf98wyu9876uhzxiou#@%%"

func Gentoken (user model.User)(string, *JwtCustomClaims, error){
	claims := &JwtCustomClaims{
		UserId: user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resut, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", nil, err
	}
	return resut, claims, nil
}

func GenRefreshtoken (user model.User)(string, *JwtCustomClaims, error){
	claims := &JwtCustomClaims{
		UserId: user.ID,
		Email: user.Email,
		Role_Ticker: user.RoleTicker,
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

func VerifyToken(accessToken string)(*JwtCustomClaims, error){
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

type JwtCustomClaims struct {
	UserId string
	Email  string
	Role_Ticker string
	jwt.StandardClaims
}