package database

import "Recruitment-Managment-system/models"

func GetAllJobs() (jobs []models.Job, err error) {

	query := `select * from jobs`
	jobs, err = ReadAllJobs(query)
	if err != nil {
		return
	}
	return
}
