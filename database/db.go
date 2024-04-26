package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDatabase() (sqlite *sql.DB, err error) {
	sqlite, err = sql.Open("sqlite", "sqlite.db")
	if err != nil {
		return
	}

	//defer sqlite.Close()
	sqlite.Ping()

	queryUser := `CREATE TABLE IF NOT EXISTS users(
		user_id TEXT,
		name TEXT,
		email TEXT,
		address TEXT,
		user_type TEXT,
		password TEXT,
		profile_headline TEXT
		
	)`
	err = CreateTable(queryUser, sqlite)
	if err != nil {
		return
	}

	queryProfile := `CREATE TABLE IF NOT EXISTS profile(
		profile_id TEXT,
		applicant TEXT,
		resume_file_address TEXT,
		skills TEXT,
		education TEXT,
		experience TEXT,
		name TEXT,
		email TEXT,
		phone TEXT,
		FOREIGN KEY(applicant) REFERENCES users(user_id)

	)`
	err = CreateTable(queryProfile, sqlite)
	if err != nil {
		return
	}

	queryJob := `CREATE TABLE IF NOT EXISTS jobs(
		job_id TEXT,
		title TEXT,
		description TEXT,
		posted_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		total_applications INTEGER,
		company_name TEXT,
		posted_by TEXT,
		FOREIGN KEY(posted_by) REFERENCES users(user_id)

	)`
	err = CreateTable(queryJob, sqlite)
	if err != nil {
		return

	}

	queryJobApplications := `CREATE TABLE IF NOT EXISTS jonApplications(
		job_id TEXT,
		user_id TEXT,
		FOREIGN KEY(job_id) REFERENCES jobs(job_id),
		FOREIGN KEY(user_id) REFERENCES users(user_id)

	)`
	err = CreateTable(queryJobApplications, sqlite)
	if err != nil {
		return

	}
	return

}
