package models

type User struct {
	UserID          string  `json:"user_id,omitempty"`
	Name            string  `json:"name"`
	Email           string  `json:"email,omitempty"`
	Address         string  `json:"address,omitempty"`
	UserType        string  `json:"user_type,omitempty"` // Admin or Applicant
	Password        string  `json:"password,omitempty"`
	ProfileHeadline string  `json:"profile_headline,omitempty"`
	Profile         Profile `json:"profile,omitempty"`
}

type Profile struct {
	ProfileID         string `json:"profile_id,omitempty"`
	Applicant         string `json:"applicant,omitempty"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	ResumeFileAddress string `json:"resume_file_address"`
	Skills            string `json:"skills"`
	Education         string `json:"education"`
	Experience        string `json:"experience"`
	Phone             string `json:"phone"`
}

type Job struct {
	JobID             string `json:"job_id,omitempty"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	PostedOn          string `json:"posted_on,omitempty"`
	TotalApplications int    `json:"total_applications"`
	CompanyName       string `json:"company_name"`
	PostedBy          string `json:"posted_by,omitempty"`
}

// ResumeResponse represents the JSON response from the resume parsing API
type ResumeResponse struct {
	Education  []Education  `json:"education"`
	Email      string       `json:"email"`
	Experience []Experience `json:"experience"`
	Name       string       `json:"name"`
	Phone      string       `json:"phone"`
	Skills     []string     `json:"skills"`
}

// Education represents education details in the resume
type Education struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Experience represents experience details in the resume
type Experience struct {
	Dates []string `json:"dates"`
	Name  string   `json:"name"`
	URL   string   `json:"url"`
}

type JobApplicationData struct {
	Job           Job  `json:"job"`
	ApplicantData User `json:"applicant_data"`
}
