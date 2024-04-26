package database

import (
	"Recruitment-Managment-system/models"
	"context"
	"database/sql"
	"time"
)

var (
	Sqlite *sql.DB
)

func CreateTable(query string, sqlite *sql.DB) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	_, err = sqlite.ExecContext(ctx, query)
	if err != nil {
		return
	}
	return

}

func Insert(query string, data ...any) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	_, err = Sqlite.QueryContext(ctx, query, data...)
	if err != nil {
		return
	}
	return
}

func UpdateOne(query string, data ...any) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	_, err = Sqlite.ExecContext(ctx, query, data...)
	if err != nil {
		return
	}
	return nil
}

func Read(query string, cond any, data ...any) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	err = Sqlite.QueryRowContext(ctx, query, cond).Scan(data...)
	if err != nil {
		return
	}
	//fmt.Print("data :", rows)
	return
}

func ReadMultipleRow(query string, cond any) (Data []models.JobApplicationData, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	rows, err := Sqlite.QueryContext(ctx, query, cond)
	if err != nil {
		return
	}
	data := models.JobApplicationData{}
	for rows.Next() {
		jobs := models.Job{}
		user := models.User{}
		err = rows.Scan(&jobs.JobID, &jobs.Title, &jobs.Description, &jobs.PostedOn, &jobs.TotalApplications, &jobs.CompanyName, &jobs.PostedBy, &user.Name, &user.Email, &user.ProfileHeadline, &user.Profile.Skills, &user.Profile.Experience, &user.Profile.Phone)
		if err != nil {
			return
		}
		data = models.JobApplicationData{
			Job:           jobs,
			ApplicantData: user,
		}
		Data = append(Data, data)
	}
	return
}

func ReadManyUsers(query string, cond string) (users []models.User, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	rows, err := Sqlite.QueryContext(ctx, query, cond)
	if err != nil {
		return
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.Name, &user.Email, &user.Address)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func ReadAllJobs(query string) (jobs []models.Job, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	rows, err := Sqlite.QueryContext(ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		job := models.Job{}
		err = rows.Scan(&job.JobID, &job.Title, &job.CompanyName, &job.Description, &job.TotalApplications, &job.PostedBy, &job.PostedOn)
		if err != nil {
			return
		}
		jobs = append(jobs, job)
	}

	return
}
