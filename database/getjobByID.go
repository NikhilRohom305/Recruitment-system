package database

import "Recruitment-Managment-system/models"

func GetJob(jobID string) (data []models.JobApplicationData, err error) {
	query := `SELECT distinct
	jobs.job_id,
	jobs.title,
	jobs.description,
	jobs.posted_on,
    jobs.total_applications,
	jobs.company_name,
	jobs.posted_by AS job_posted_by,
	users.name,
    users.email,
	users.profile_headline,
	profile.skills,
	profile.experience,
	profile.phone


FROM
	jobs
JOIN
	 jobApplications on jobApplications.job_id=jobs.job_id
JOIN
	users on users.user_id=jobApplications.user_id
JOIN
 profile on profile.applicant=jobApplications.user_id
WHERE
   jobs.job_id =?
`

	data, err = ReadMultipleRow(query, jobID)
	if err != nil {
		return
	}
	return
}
