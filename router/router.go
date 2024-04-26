package router

import (
	"Recruitment-Managment-system/controllers"
	"Recruitment-Managment-system/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) (err error) {

	router.HandleFunc("/signup", controllers.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	router.HandleFunc("/uploadResume", middleware.ApplicantAuth(controllers.UploadResume)).Methods(http.MethodPost)
	router.HandleFunc("/admin/job", middleware.AdminAuth((controllers.CreateJob))).Methods(http.MethodPost)

	router.HandleFunc("/admin/job/{job_id}", middleware.AdminAuth(controllers.GetJobByJobID)).Methods(http.MethodGet)

	router.HandleFunc("/jobs/apply", middleware.ApplicantAuth(controllers.Apply)).Methods(http.MethodGet)

	router.HandleFunc("/admin/applicants", middleware.AdminAuth(controllers.GetAllApplicants)).Methods(http.MethodGet)

	router.HandleFunc("/jobs", controllers.GetAllJobOpenings).Methods(http.MethodGet)

	router.HandleFunc("/admin/applicant/{applicant_id}", middleware.AdminAuth(controllers.GetApplicant)).Methods(http.MethodGet)
	return
}
