package controllers

import (
	"Recruitment-Managment-system/models"
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreateJob(w http.ResponseWriter, req *http.Request) {
	role := req.Context().Value("role").(string)
	userID := req.Context().Value("userID").(string)

	if role != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	job := new(models.Job)
	job.PostedBy = userID
	err := json.NewDecoder(req.Body).Decode(&job)
	if err != nil {
		return
	}
	job.JobID = uuid.NewString()
	err = services.CreateJob(job)
	if err != nil {
		return
	}
	resp := make(map[string]interface{})
	resp["msg"] = "success"
	resp["job_id"] = job.JobID
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}
