package database

import "Recruitment-Managment-system/models"

func InsertUser(user *models.User) (err error) {
	query := `INSERT INTO users (user_id,name,email,address,user_type,password,profile_headline) VALUES (?,?,?,?,?,?,?)`
	err = Insert(query, &user.UserID, &user.Name, &user.Email, &user.Address, &user.UserType, &user.Password, &user.ProfileHeadline)
	if err != nil {
		return
	}

	queryProfile := `INSERT INTO profile(profile_id,applicant,resume_file_address,skills,education,experience,name,email,phone) VALUES (?,?,?,?,?,?,?,?,?)`
	err = Insert(queryProfile, &user.Profile.ProfileID, &user.UserID, &user.Profile.ResumeFileAddress, &user.Profile.Skills, &user.Profile.Education, &user.Profile.Experience, &user.Name, &user.Email, &user.Profile.Phone)
	if err != nil {
		return
	}
	return
}
