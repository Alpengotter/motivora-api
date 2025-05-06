package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CompanyRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", GetCompaniesHandler).Methods("GET")

	r.HandleFunc("/get-all-stat", GetCompaniesStatisticHandler).Methods("GET")
	return r
}

func GetCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)

	companies, err := GetCompanies(offset, limit)
	if err != nil {
		http.Error(w, "Error fetching companies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(companies)
}

func GetCompaniesStatisticHandler(w http.ResponseWriter, r *http.Request) {
	stat, err := GetCompaniesStatistic()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stat)
}
