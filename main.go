package main

import (
	"log"
	"net/http"

	"github.com/athomas5/go-rest/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	h := handlers.NewApiHandler()
	h.InitCompanies()

	r.HandleFunc("/api/companies", h.GetCompanies).Methods("GET")
	r.HandleFunc("/api/companies/{id}", h.GetCompany).Methods("GET")
	r.HandleFunc("/api/companies", h.CreateCompany).Methods("POST")
	r.HandleFunc("/api/companies/{id}", h.UpdateCompany).Methods("PUT")
	r.HandleFunc("/api/companies/{id}", h.DeleteCompany).Methods("DELETE")

	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
