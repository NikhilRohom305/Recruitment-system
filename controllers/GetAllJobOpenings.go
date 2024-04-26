package controllers

import (
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"
)

func GetAllJobOpenings(w http.ResponseWriter, req *http.Request) {

	jobs, err := services.GetAllJobOpenings()
	if err != nil {
		return
	}

	resp := make(map[string]interface{})
	resp["jobs"] = jobs
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

}
