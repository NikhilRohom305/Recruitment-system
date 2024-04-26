package services

import "Recruitment-Managment-system/database"

func Application(jobID string, userID string) (err error) {
	err = database.Apply(jobID, userID)
	if err != nil {
		return
	}
	return
}
