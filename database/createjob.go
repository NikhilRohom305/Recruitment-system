package database

import (
	"Recruitment-Managment-system/models"
)

func CreateJob(job *models.Job) (err error) {

	query := `INSERT INTO jobs(job_id,title,description,posted_on,total_applications,company_name,posted_by) VALUES(?,?,?,?,?,?,?)`
	err = Insert(query, &job.JobID, &job.Title, &job.Description, &job.PostedOn, &job.TotalApplications, &job.CompanyName, &job.PostedBy)
	if err != nil {
		return
	}
	return
}
