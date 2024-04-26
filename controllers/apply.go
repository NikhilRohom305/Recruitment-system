package controllers

import (
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"
)

func Apply(w http.ResponseWriter, req *http.Request) {
	jobID := req.URL.Query().Get("job_id")

	userID := req.Context().Value("userID").(string)

	err := services.Application(jobID, userID)
	if err != nil {
		http.Error(w, "error in application", 400)
	}
	resp := make(map[string]interface{})
	resp["msg"] = "success"
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}

}
