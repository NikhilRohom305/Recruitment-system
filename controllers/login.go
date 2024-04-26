package controllers

import (
	"Recruitment-Managment-system/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

func Login(w http.ResponseWriter, req *http.Request) {

	payload := make(map[string]interface{})
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return
	}
	email := payload["email"].(string)
	password := payload["password"].(string)
	if email == "" || password == "" {
		http.Error(w, fmt.Sprintf("\"%v\"", `Invalid request body`), http.StatusBadRequest)
		return
	}
	if err = validateEmail(email); err != nil {
		http.Error(w, fmt.Sprintf("\"%v\"", `invalid mail`), http.StatusBadRequest)
		return
	}
	token, err := services.Login(email, password)
	if err != nil {
		return
	}
	resp := make(map[string]interface{})
	resp["token"] = token
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("\"%v\"", `error in login`), http.StatusBadRequest)

	}

}

func validateEmail(email string) (err error) {
	em := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	if !em.MatchString(email) {
		err = errors.New("invalid email")
		return
	}
	return
}
