package services

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/models"
)

func GetJobByID(jobID string) (data []models.JobApplicationData, err error) {

	data, err = database.GetJob(jobID)
	if err != nil {
		return
	}
	return
}
