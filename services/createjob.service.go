package services

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/models"
)

func CreateJob(job *models.Job) (err error) {
	err = database.CreateJob(job)
	if err != nil {
		return
	}

	return
}
