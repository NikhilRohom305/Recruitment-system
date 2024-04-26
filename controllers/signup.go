package controllers

import (
	"Recruitment-Managment-system/models"
	"Recruitment-Managment-system/services"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, req *http.Request) {
	user := new(models.User)
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid payload", 500)
		return
	}
	if user.Name != "" && user.Email != "" {
		err = services.CreateUser(user)
		if err != nil {
			return
		}
	}

	response := make(map[string]interface{})
	response["message"] = "success"
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "error in registration", 404)
	}

}
