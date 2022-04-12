package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var companies []Company

type Company struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Rating   int       `json:"rating"`
	Features *Features `json:"features"`
}

type Features struct {
	Pros string `json:"pros"`
	Cons string `json:"const"`
}

type ApiHandler struct {
}

func (h *ApiHandler) InitCompanies() {
	companies = append(companies, Company{ID: "1", Name: "Google", Rating: 4, Features: &Features{Pros: "Prestige, WLB, Culture, Office", Cons: "Angular, Salary"}})
	companies = append(companies, Company{ID: "2", Name: "Stripe", Rating: 5, Features: &Features{Pros: "React, Salary, Office", Cons: "WLB, Culture"}})
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

func (h *ApiHandler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
	log.Println("Fetching companies")
}

func (h *ApiHandler) GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, company := range companies {
		if company.ID == params["id"] {
			json.NewEncoder(w).Encode(company)
			log.Printf("Fetching company %v\n", params["id"])
			return
		}
	}

	log.Printf("Fetching company %v failed\n", params["id"])
	json.NewEncoder(w).Encode("Fetching company failed")
}

func (h *ApiHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, company := range companies {
		if company.ID == params["id"] {
			companies = append(companies[:i], companies[i+1:]...)
			log.Printf("Deleting company %v\n", params["id"])
			json.NewEncoder(w).Encode(companies)
			return
		}
	}

	errMessage, _ := fmt.Printf("Deleting company %v failed\n", params["id"])
	json.NewEncoder(w).Encode(errMessage)
}

func (h *ApiHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var newCompany Company
	_ = json.NewDecoder(r.Body).Decode(&newCompany)
	newCompany.ID = params["id"]

	for i, company := range companies {
		if company.ID == params["id"] {
			companies[i] = newCompany
			log.Printf("Updated company %v\n", params["id"])
			json.NewEncoder(w).Encode(newCompany)
			return
		}
	}

	log.Printf("Updateding company %v failed\n", params["id"])
	json.NewEncoder(w).Encode("Updateding company failed")
}

func (h *ApiHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var company Company
	_ = json.NewDecoder(r.Body).Decode(&company)
	company.ID = strconv.Itoa(rand.Intn(10000000))
	companies = append(companies, company)

	log.Println("New company was created")
	json.NewEncoder(w).Encode(company)
}
