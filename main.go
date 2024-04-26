package main

import (
	"Recruitment-Managment-system/database"
	"Recruitment-Managment-system/router"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	route := new(mux.Router)
	err := router.SetupRoutes(route)
	if err != nil {
		return
	}
	database.Sqlite, err = database.InitDatabase()
	if err != nil {
		return
	}

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:8080"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowCredentials(),
	)(route)
	err = http.ListenAndServe(":8080", cors)
	if err != nil {
		return
	}

}
