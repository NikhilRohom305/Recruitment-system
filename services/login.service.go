package services

import (
	"Recruitment-Managment-system/database"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var SecretKey = []byte("81mohomlihkin")

func Login(email string, password string) (token string, err error) {

	Password := HashPassword(password)
	userID, userType, err := database.LoginUser(email, Password)
	if err != nil {
		return
	}
	token, err = GenerateToken(userType, userID, email)
	if err != nil {
		logrus.WithField("err", err.Error()).Error("error generating jwt token for user")
		return

	}

	return
}

func GenerateToken(role string, UserID string, email string) (token string, err error) {
	tokenExpirationTime := time.Now().Add(time.Hour * 24)
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Role":    role,
		"user_id": UserID,
		"email":   email,
		"exp":     tokenExpirationTime.Unix(),
	})
	token, err = tokenObject.SignedString(SecretKey)
	return
}
