package router

import (
	"github.com/chenshanmugarajah/chens-job-portal-api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/jobs", controller.GetMyAllJobs).Methods("GET")
	router.HandleFunc("/api/job/{id}", controller.GetOneJobOnly).Methods("GET")
	router.HandleFunc("/api/job", controller.CreateJob).Methods("POST")
	router.HandleFunc("/api/job/{id}", controller.UpdateOneJobOnly).Methods("PUT")
	router.HandleFunc("/api/job/{id}", controller.DeleteAJob).Methods("DELETE")
	router.HandleFunc("/api/deletealljob", controller.DeleteAllJobs).Methods("DELETE")

	return router
}
