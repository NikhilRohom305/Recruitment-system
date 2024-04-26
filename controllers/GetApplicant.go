package controllers

import (
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetApplicant(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	applicantID := vars["applicant_id"]

	applicantData, err := services.GetApplicantData(applicantID)
	if err != nil {
		return
	}
	resp := make(map[string]interface{})
	resp["applicant_Data"] = applicantData
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "cant get applicant data", http.StatusBadRequest)
		return
	}
}
