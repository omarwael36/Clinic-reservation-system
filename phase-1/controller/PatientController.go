package controller

import (
	"Clinic-Reservation-System/config"
	"Clinic-Reservation-System/helper"
	"Clinic-Reservation-System/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PatientSignUp(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPost {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.Patient
	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		log.Printf("Error binding form data: %v", err)
		return
	}

	if input.Name == "" || input.Email == "" || input.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE patientEmail = ?", input.Email).Scan(&count)
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User Already Exist"})
		return
	}

	_, err := db.Exec("INSERT INTO patient (patientName, patientEmail, patientPassword) VALUES (?, ?, ?)",
		input.Name, input.Email, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error inserting data into the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Sign-up is done successfully!"
	log.Print("Data inserted into the database")

	c.JSON(http.StatusOK, response)
}

func PatientSignIn(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPost {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.Patient
	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		log.Printf("Error binding form data: %v", err)
		return
	}

	if input.Email == "" || input.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE patientEmail = ?", input.Email).Scan(&count)
	if count == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "There is no such user!"})
		return
	}

	var storedPassword string
	var PatientID int
	err := db.QueryRow("SELECT patientPassword, PatientID FROM patient WHERE patientEmail = ?", input.Email).Scan(&storedPassword, &PatientID)

	if err != nil || input.Password != storedPassword {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid Credentials"})
		log.Printf("Error Sign in: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Sign-in is done successfully!"
	log.Print("Sign-in is done successfully!")
	helper.StoreUserInSession(c.Writer, c.Request, PatientID, "Patient")
	c.JSON(http.StatusOK, response)
}

func ShowAllDoctors(c *gin.Context) {
	var doctor model.DoctorResponse
	var response model.Response
	var arrDoctorResponse []model.DoctorResponse

	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodGet {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
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

	c.JSON(http.StatusOK, response)
}

func ShowDoctorSlots(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodGet {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	doctorID := c.Query("DoctorID")

	var doctorCount int
	err := db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorID = ?", doctorID).Scan(&doctorCount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error querying doctor count: %v", err)
		return
	}

	if doctorCount == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Doctor with provided DoctorID does not exist"})
		return
	}

	rows, err := db.Query("SELECT SlotID, SlotDateTime, DoctorID FROM slot WHERE DoctorID = ? AND PatientID IS NULL", doctorID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error selecting data from the database: %v", err)
		return
	}

	var slots []model.Slot
	for rows.Next() {
		var slot model.Slot

		err := rows.Scan(&slot.SlotID, &slot.SlotDateTime, &slot.DoctorID)
		if err != nil {
			log.Printf("Error scanning slot: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Append the slot to the slots slice
		slots = append(slots, slot)
	}

	response.Status = http.StatusOK
	response.Message = "Slots retrieved successfully!"
	response.Data = slots

	c.JSON(http.StatusOK, response)
}

func ReserveSlot(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPut {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.ReserveSlotRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		log.Printf("Error binding form data: %v", err)
		return
	}
	patientID, userType := helper.GetUserFromSession(c.Request)
	if userType != "Patient" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var slotCount int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE SlotID = ?", input.SlotID).Scan(&slotCount)
	if slotCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slot with provided SlotID does not exist"})
		return
	}

	var patientCount int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE PatientID = ?", patientID).Scan(&patientCount)
	if patientCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient with provided PatientID does not exist"})
		return
	}

	_, err := db.Exec("UPDATE slot SET PatientID = ? WHERE SlotID = ?", patientID, input.SlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error updating slot in the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Slot reserved successfully!"
	log.Print("Slot reserved in the database")

	c.JSON(http.StatusOK, response)
}

func UpdateAppointment(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPut {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.UpdateAppointmentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		log.Printf("Error binding form data: %v", err)
		return
	}

	patientID, userType := helper.GetUserFromSession(c.Request)
	if userType != "Patient" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var appointmentCount int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE SlotID = ?", input.AppointmentID).Scan(&appointmentCount)
	if appointmentCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment with provided AppointmentID does not exist"})
		return
	}

	var slotCount int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE SlotID = ? AND DoctorID = ?", input.NewSlotID, input.NewDoctorID).Scan(&slotCount)
	if slotCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slot with provided NewSlotID and NewDoctorID does not exist"})
		return
	}

	var patientCount int
	db.QueryRow("SELECT COUNT(*) FROM patient WHERE PatientID = ?", patientID).Scan(&patientCount)
	if patientCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient with provided NewPatientID does not exist"})
		return
	}

	_, err := db.Exec("UPDATE slot SET PatientID = NULL WHERE SlotID = ?", input.AppointmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error updating appointment in the database: %v", err)
		return
	}

	_, err = db.Exec("UPDATE slot SET PatientID = ? WHERE SlotID = ?", patientID, input.NewSlotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error updating appointment in the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Appointment updated successfully!"
	log.Print("Appointment updated in the database")

	c.JSON(http.StatusOK, response)
}

func CancelAppointment(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodDelete {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	patientID, userType := helper.GetUserFromSession(c.Request)
	if userType != "Patient" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	slotIDStr := c.Query("slotId")

	if slotIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slot ID is required"})
		return
	}

	slotID, err := strconv.Atoi(slotIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid slot ID"})
		log.Printf("Error parsing slot ID: %v", err)
		return
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE PatientID = ? AND SlotID = ?", patientID, slotID).Scan(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No appointment with this ID"})
		return
	}

	_, err = db.Exec("UPDATE slot SET PatientID = NULL WHERE PatientID = ? AND SlotID = ?", patientID, slotID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Cannot cancel appointment: %v", err)
		return
	}

	log.Print("Cancel appointment success")

	response.Status = http.StatusOK
	response.Message = "Cancel appointment success"
	log.Print("Appointment canceled from the database")

	c.JSON(http.StatusOK, response)
}

func ShowAllReservations(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	patientID, userType := helper.GetUserFromSession(c.Request)
	if userType != "Patient" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM slot WHERE PatientID = ? ", patientID).Scan(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No reserved appointments"})
		return
	}

	rows, err := db.Query("SELECT SlotID, SlotDateTime, DoctorID, PatientID FROM slot WHERE PatientID = ? ", patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
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

	c.JSON(http.StatusOK, response)
}
