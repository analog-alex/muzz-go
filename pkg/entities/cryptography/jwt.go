package cryptography

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// TODO move to configuration file
const superDuperSecret = "secret_singing_key"

// GenerateJWToken create a jwt token with the given id
func GenerateJWToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(superDuperSecret))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}
	return tokenString, nil
}

// VerifyJWToken verifies a jwt token
func VerifyJWToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(superDuperSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		validator := jwt.NewValidator()

		// additional validations including expiration time and not before time
		err := validator.Validate(claims)
		if err != nil {
			return "", errors.New("token is invalid")
		}

		// get claims sub as string
		return claims["sub"].(string), nil
	} else {
		return "", err
	}
}
