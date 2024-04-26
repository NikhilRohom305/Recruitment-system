package controllers

import (
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"
)

func GetAllApplicants(w http.ResponseWriter, req *http.Request) {

	applicants, err := services.GetAllApplicants()
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	resp := make(map[string]interface{})
	resp["users"] = applicants
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
	}

}
