package controller

import (
	"Clinic-Reservation-System/config"
	"Clinic-Reservation-System/helper"
	"Clinic-Reservation-System/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func PatientSignUp(w http.ResponseWriter, r *http.Request) {
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

	name := r.FormValue("patientName")
	email := r.FormValue("patientEmail")
	password := r.FormValue("patientPassword")

	if name == "" || email == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	var count int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE patientEmail = ?", email).Scan(&count)
	if count > 0 {
		http.Error(w, "User Already Exist", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO patient (patientName, patientEmail, patientPassword) VALUES (?, ?, ?)", name, email, password)
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

func PatientSignIn(w http.ResponseWriter, r *http.Request) {
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
	email := r.FormValue("patientEmail")
	password := r.FormValue("patientPassword")
	if email == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	var count int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE patientEmail = ?", email).Scan(&count)
	if count == 0 {
		http.Error(w, "There is no such user!", http.StatusBadRequest)
		return
	}
	var storedPassword string
	var PatientID int
	err = db.QueryRow("SELECT patientPassword, PatientID FROM patient WHERE patientEmail = ?", email).Scan(&storedPassword, &PatientID)

	if err != nil || password != storedPassword {
		http.Error(w, "Invalid Credintials", http.StatusInternalServerError)
		log.Printf("Error Sign in: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Sign-in is done successfully!"
	log.Print("Sign-in is done successfully!")
	helper.StoreUserInSession(w, r, PatientID, "Patient")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ShowAllDoctors(w http.ResponseWriter, r *http.Request) {
	var doctor model.DoctorResponse
	var response model.Response
	var arrDoctorResponse []model.DoctorResponse

	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT DoctorId, DoctorName, DoctorEmail FROM doctor")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.Email)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrDoctorResponse = append(arrDoctorResponse, doctor)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDoctorResponse

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func ShowDoctorSlots(w http.ResponseWriter, r *http.Request) {
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
	doctorID := r.FormValue("DoctorID")
	var doctorCount int
	err = db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorID = ?", doctorID).Scan(&doctorCount)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error querying doctor count: %v", err)
		return
	}

	if doctorCount == 0 {
		http.Error(w, "Doctor with provided DoctorID does not exist", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT SlotID, SlotDateTime, DoctorID FROM slot WHERE DoctorID = ? AND PatientID IS NULL", doctorID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error selecting data from the database: %v", err)
		return
	}

	var slots []model.Slot
	for rows.Next() {
		var slot model.Slot

		err := rows.Scan(&slot.SlotID, &slot.SlotDateTime, &slot.DoctorID)
		if err != nil {
			log.Printf("Error scanning slot: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Append the slot to the slots slice
		slots = append(slots, slot)
	}

	response.Status = http.StatusOK
	response.Message = "Slots retrieved successfully!"
	response.Data = slots

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ReserveSlot(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("Error parsing form data: %v", err)
		return
	}

	SlotIDStr := r.FormValue("SlotID")
	PatientID, _ := helper.GetUserFromSession(r)

	SlotID, err := strconv.Atoi(SlotIDStr)
	if err != nil {
		http.Error(w, "Invalid SlotID", http.StatusBadRequest)
		log.Printf("Error parsing SlotID: %v", err)
		return
	}

	var slotCount int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE SlotID = ?", SlotID).Scan(&slotCount)
	if slotCount == 0 {
		http.Error(w, "Slot with provided SlotID does not exist", http.StatusBadRequest)
		return
	}

	var patientCount int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE PatientID = ?", PatientID).Scan(&patientCount)
	if patientCount == 0 {
		http.Error(w, "Patient with provided PatientID does not exist", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE slot SET PatientID = ? WHERE SlotID = ?", PatientID, SlotID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error updating slot in the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Slot reserved successfully!"
	log.Print("Slot reserved in the database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		log.Printf("Error parsing form data: %v", err)
		return
	}

	AppointmentIDStr := r.FormValue("AppointmentID")
	NewSlotIDStr := r.FormValue("NewSlotID")
	NewDoctorIDStr := r.FormValue("NewDoctorID")
	NewPatientIDStr := r.FormValue("NewPatientID")

	AppointmentID, err := strconv.Atoi(AppointmentIDStr)
	if err != nil {
		http.Error(w, "Invalid AppointmentID", http.StatusBadRequest)
		log.Printf("Error parsing AppointmentID: %v", err)
		return
	}

	NewSlotID, err := strconv.Atoi(NewSlotIDStr)
	if err != nil {
		http.Error(w, "Invalid NewSlotID", http.StatusBadRequest)
		log.Printf("Error parsing NewSlotID: %v", err)
		return
	}

	NewDoctorID, err := strconv.Atoi(NewDoctorIDStr)
	if err != nil {
		http.Error(w, "Invalid NewDoctorID", http.StatusBadRequest)
		log.Printf("Error parsing NewDoctorID: %v", err)
		return
	}

	NewPatientID, err := strconv.Atoi(NewPatientIDStr)
	if err != nil {
		http.Error(w, "Invalid NewPatientID", http.StatusBadRequest)
		log.Printf("Error parsing NewPatientID: %v", err)
		return
	}
	var appointmentCount int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE SlotID = ?", AppointmentID).Scan(&appointmentCount)
	if appointmentCount == 0 {
		http.Error(w, "Appointment with provided AppointmentID does not exist", http.StatusBadRequest)
		return
	}
	var slotCount int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE SlotID = ? AND DoctorID = ?", NewSlotID, NewDoctorID).Scan(&slotCount)
	if slotCount == 0 {
		http.Error(w, "Slot with provided NewSlotID and NewDoctorID does not exist", http.StatusBadRequest)
		return
	}
	var patientCount int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE PatientID = ?", NewPatientID).Scan(&patientCount)
	if patientCount == 0 {
		http.Error(w, "Patient with provided NewPatientID does not exist", http.StatusBadRequest)
		return
	}
	_, err = db.Exec("UPDATE slot SET PatientID = NULL WHERE SlotID = ?", AppointmentID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error updating appointment in the database: %v", err)
		return
	}

	_, err = db.Exec("UPDATE slot SET PatientID = ? WHERE SlotID = ?", NewPatientID, NewSlotID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Error updating appointment in the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Appointment updated successfully!"
	log.Print("Appointment updated in the database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CancelAppointment(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	patientID, _ := helper.GetUserFromSession(r)

	slotIDStr := r.FormValue("slotId")

	if slotIDStr == "" {
		http.Error(w, "Slot ID is required", http.StatusBadRequest)
		return
	}

	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		http.Error(w, "Invalid slot ID", http.StatusBadRequest)
		log.Printf("Error parsing slot ID: %v", err)
		return
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE PatientID = ? AND SlotID = ?", patientID, slotID).Scan(&count)
	if count == 0 {
		http.Error(w, "No appointment with this ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE slot SET PatientID = NULL WHERE PatientID = ? AND SlotID = ?", patientID, slotID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Cannot cancel appointment: %v", err)
		return
	}

	log.Print("Cancel appointment success")

	response.Status = http.StatusOK
	response.Message = "Cancel appointment success"
	log.Print("Appointment canceled from the database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ShowAllReservations(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	patientID, _ := helper.GetUserFromSession(r)

	var count int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE PatientID = ? ", patientID).Scan(&count)
	if count == 0 {
		http.Error(w, "No reserved appointments", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT SlotID, SlotDateTime, DoctorID, PatientID FROM slot WHERE PatientID = ? ", patientID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Cannot show reservations: %v", err)
		return
	}
	defer rows.Close()

	var reservedSlots []model.Slot

	for rows.Next() {
		var slot model.Slot
		err := rows.Scan(&slot.SlotID, &slot.SlotDateTime, &slot.DoctorID, &slot.PatientID)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			// Parse the SlotDateTime using the new method
			parsedTime, err := slot.ParseSlotDateTime()
			if err != nil {
				log.Fatal(err.Error())
			}
			// Assign the parsed time back to the struct
			slot.SlotDateTime = parsedTime.String()
			reservedSlots = append(reservedSlots, slot)
		}
	}

	log.Print("Printed all reservations")

	response.Status = http.StatusOK
	response.Message = "Printed all reservations"
	response.Data = reservedSlots

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
