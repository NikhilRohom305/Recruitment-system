package services

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/models"
)

func GetAllJobOpenings() (jobs []models.Job, err error) {
	jobs, err = database.GetAllJobs()
	if err != nil {
		return
	}
	return
}
