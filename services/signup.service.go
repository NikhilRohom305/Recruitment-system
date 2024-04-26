package services

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/models"
	"crypto/sha256"
	"encoding/base64"

	"github.com/google/uuid"
)

func CreateUser(user *models.User) (err error) {
	user.Password = HashPassword(user.Password)
	user.UserID = uuid.NewString()

	err = database.InsertUser(user)
	if err != nil {
		return
	}
	return
}

func HashPassword(password string) string {

	hsha := sha256.New()
	hsha.Write([]byte(password))
	hash := base64.URLEncoding.EncodeToString(hsha.Sum(nil))
	//logrus.Info(password, " -> ", hash)
	return hash
}
