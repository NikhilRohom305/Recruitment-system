package controllers

import (
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetJobByJobID(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	jobID := vars["job_id"]

	data, err := services.GetJobByID(jobID)
	if err != nil {
		http.Error(w, "job id not found", http.StatusBadRequest)
	}
	resp := make(map[string]interface{})
	resp["msg"] = "success"
	resp["applicants_Data"] = data
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "cannot get job info", http.StatusBadRequest)
	}

}
