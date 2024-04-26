package services

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/models"
)

func GetAllApplicants() (applicants []models.User, err error) {
	applicants, err = database.GetApplicants()
	if err != nil {
		return
	}

	return
}
