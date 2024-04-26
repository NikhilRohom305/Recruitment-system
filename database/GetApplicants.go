package database

import "Recruitment-Managment-system/models"

func GetApplicants() (data []models.User, err error) {
	usertype := "Applicant"
	query := `select * from users where user_type=? `
	data, err = ReadManyUsers(query, usertype)
	if err != nil {
		return
	}
	return
}
