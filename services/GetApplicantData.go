package services

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/models"
)

func GetApplicantData(applicantID string) (user models.Profile, err error) {

	user, err = database.GetApplicantData(applicantID)
	if err != nil {
		return
	}
	return
}
