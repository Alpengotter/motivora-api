package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"motivora-backend/internal/models"
	"net/http"
	"strconv"
)

// EmployerRouter возвращает http.HandlerFunc, который обрабатывает маршруты /employers/
func EmployerRouter() http.Handler {
	r := mux.NewRouter()

	// Здесь "/" — это базовый путь после StripPrefix
	r.HandleFunc("/", GetEmployersHandler).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", GetEmployerByIDHandler).Methods("GET")
	r.HandleFunc("/", CreateEmployerHandler).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", UpdateEmployerHandler).Methods("PUT")

	r.HandleFunc("/get-all-stat", GetEmployerStatisticHandler).Methods("GET")

	return r
}

func GetEmployersHandler(w http.ResponseWriter, r *http.Request) {
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)

	// Вызов GetUsers из db
	user, err := GetUsers(offset, limit)
	if err != nil {
		http.Error(w, "Error fetching employers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}

func GetEmployerByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employer ID", http.StatusBadRequest)
		return
	}

	employer, err := GetUserByID(id)
	if err != nil {
		http.Error(w, "Error fetching employer", http.StatusInternalServerError)
		return
	}

	if employer == nil {
		http.Error(w, "Employer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(employer)
}

func CreateEmployerHandler(w http.ResponseWriter, r *http.Request) {
	var employer models.User
	err := json.NewDecoder(r.Body).Decode(&employer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := CreateUser(employer)
	if err != nil {
		http.Error(w, "Error creating employer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func UpdateEmployerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employer ID", http.StatusBadRequest)
		return
	}

	var updateReq models.UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = UpdateUser(id, updateReq)
	if err != nil {
		http.Error(w, "Error updating employer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Employer updated successfully",
	})
}

func GetEmployerStatisticHandler(w http.ResponseWriter, r *http.Request) {
	stat, err := GetEmployerStatistic()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stat)
}
