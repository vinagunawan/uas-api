package handlers

import (
	"backend/core"
	"backend/service/entities"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Login handler
func Login(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	var dbUser entities.User
	core.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser)
	if dbUser.ID != 0 {
		json.NewEncoder(w).Encode(dbUser)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// Get user data handler
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user entities.User
	core.DB.First(&user, userId)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Update user data handler
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user entities.User
	core.DB.First(&user, userId)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var updatedUser entities.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	user.Weight = updatedUser.Weight
	user.Height = updatedUser.Height
	user.Age = updatedUser.Age

	core.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

// Record water intake handler
func RecordWaterIntake(w http.ResponseWriter, r *http.Request) {
	var intake entities.WaterIntake
	json.NewDecoder(r.Body).Decode(&intake)
	intake.Timestamp = time.Now()

	core.DB.Create(&intake)
	json.NewEncoder(w).Encode(intake)
}

// Get water intake history handler
func GetWaterIntake(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var intakes []entities.WaterIntake
	core.DB.Where("user_id = ?", userId).Find(&intakes)
	json.NewEncoder(w).Encode(intakes)
}
