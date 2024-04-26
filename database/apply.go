package database

func Apply(jobID string, userID string) (err error) {
	query := `INSERT INTO jobApplications (job_id,user_id) VALUES (?,?)`
	err = Insert(query, &jobID, &userID)
	if err != nil {
		return
	}
	count, err := getNoOfApplication(jobID)
	if err != nil {
		return
	}
	count += 1
	updateJob := `UPDATE jobs SET total_applications=? WHERE job_id=?`
	err = UpdateOne(updateJob, count, jobID)
	if err != nil {
		return
	}
	return
}

func getNoOfApplication(jobID string) (count int, err error) {

	query := `SELECT total_applications FROM jobs WHERE job_id=?`
	err = Read(query, jobID, &count)
	if err != nil {
		return
	}

	return
}
