package database

import "Recruitment-Managment-system/models"

func GetApplicantData(applicantID string) (user models.Profile, err error) {

	query := `select profile_id,applicant,resume_file_address,name,email,skills,education,experience,phone from  profile where applicant=?`
	err = Read(query, applicantID, &user.ProfileID, &user.Applicant, &user.ResumeFileAddress, &user.Name, &user.Email, &user.Skills, &user.Education, &user.Experience, &user.Phone)
	if err != nil {
		return
	}
	return

}
