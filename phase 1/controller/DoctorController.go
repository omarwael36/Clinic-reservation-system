package controller

import (
	"Clinic-Reservation-System/config"
	"Clinic-Reservation-System/helper"
	"Clinic-Reservation-System/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func DoctorSignUp(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("Error parsing form data: %v", err)
		return
	}

	name := r.FormValue("DoctorName")
	email := r.FormValue("DoctorEmail")
	password := r.FormValue("DoctorPassword")

	if name == "" || email == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	var count int
	db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorEmail = ?", email).Scan(&count)
	if count > 0 {
		http.Error(w, "User Already Exist", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO doctor (DoctorName, DoctorEmail, DoctorPassword) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error inserting data into the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Sign-up is done successfully!"
	log.Print("Data inserted into the database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DoctorSignIn(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("Error parsing form data: %v", err)
		return
	}
	email := r.FormValue("DoctorEmail")
	password := r.FormValue("DoctorPassword")
	if email == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	var count int
	db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorEmail = ?", email).Scan(&count)
	if count == 0 {
		http.Error(w, "There is no such user!", http.StatusBadRequest)
		return
	}
	var storedPassword string
	var DoctorID int
	err = db.QueryRow("SELECT DoctorPassword, DoctorID FROM doctor WHERE DoctorEmail = ?", email).Scan(&storedPassword, &DoctorID)

	if err != nil || password != storedPassword {
		http.Error(w, "Invalid Credintials", http.StatusInternalServerError)
		log.Printf("Error Sign in: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Sign-in is done successfully!"
	log.Print("Sign-in is done successfully!")
	helper.StoreUserInSession(w, r, DoctorID, "Doctor")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func SetSchedule(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("Error parsing form data: %v", err)
		return
	}
	DoctorID, _ := helper.GetUserFromSession(r)
	slotDateTimeStr := r.FormValue("slotDateTime")

	slotDateTime, err := time.Parse("2006-01-02T15:04:05", slotDateTimeStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		log.Printf("Error parsing slotDateTime: %v", err)
		return
	}
	log.Printf("Retrieved DoctorID from session: %d", DoctorID)
	var doctorCount int
	db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorID = ?", DoctorID).Scan(&doctorCount)
	if doctorCount == 0 {
		http.Error(w, "Doctor with provided DoctorID does not exist", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO slot (SlotDateTime, DoctorID) VALUES (?, ?)", slotDateTime, DoctorID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error inserting data into the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Slot added successfully!"
	log.Print("Slot added to the database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
