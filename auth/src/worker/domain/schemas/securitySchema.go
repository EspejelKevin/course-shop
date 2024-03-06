package schemas

import (
	"auth/src/shared/infrastructure"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var settings = infrastructure.NewSettings()

func CreateAccessToken(user map[string]interface{}) (string, error) {
	secretKey := []byte(settings.SecretKey)
	minutes, _ := strconv.Atoi(settings.TimeExpiration)
	timeExpiration := time.Duration(minutes) * time.Minute
	tokenExpiration := time.Now().Add(timeExpiration).Unix()

	claims := jwt.MapClaims{
		"exp": tokenExpiration,
		"sub": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ValidateAccessToken(tokenFromRequest string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenFromRequest, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(settings.SecretKey), nil
	})

	if err != nil {
		return map[string]interface{}{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		user := claims["sub"].(map[string]interface{})
		return user, nil
	}

	return map[string]interface{}{}, fmt.Errorf("invalid token or token has expired")
}
