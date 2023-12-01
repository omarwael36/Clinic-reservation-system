package controller

import (
	"Clinic-Reservation-System/config"
	"time"
	//"Clinic-Reservation-System/helper"
	"Clinic-Reservation-System/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func DoctorSignUp(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPost {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.Doctor
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
	db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorEmail = ?", input.Email).Scan(&count)
	if count > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User Already Exist"})
		return
	}

	_, err := db.Exec("INSERT INTO doctor (DoctorName, DoctorEmail, DoctorPassword) VALUES (?, ?, ?)",
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

func DoctorSignIn(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPost {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.Doctor
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
	db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorEmail = ?", input.Email).Scan(&count)
	if count == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "There is no such user!"})
		return
	}

	var storedPassword string
	var DoctorID int
	err := db.QueryRow("SELECT DoctorPassword, DoctorID FROM doctor WHERE DoctorEmail = ?", input.Email).Scan(&storedPassword, &DoctorID)

	if err != nil || input.Password != storedPassword {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid Credentials"})
		log.Printf("Error Sign in: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Sign-in is done successfully!"
	log.Print("Sign-in is done successfully!")
	response.Data = map[string]interface{}{"DoctorID": DoctorID} // Include PatientID in Data field
	c.JSON(http.StatusOK, response)
}

func SetSchedule(c *gin.Context) {
	var response model.Response
	db := config.DatabaseConnection()
	defer db.Close()

	if c.Request.Method != http.MethodPost {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var input model.Slot
	if err := c.ShouldBind(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		log.Printf("Error binding form data: %v", err)
		return
	}

	DoctorIDStr := c.Param("id")
	doctorID, err := strconv.Atoi(DoctorIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Doctor ID"})
		return
	}

	slotDateTime := input.SlotDateTime // Assuming it's a string from the frontend

	log.Printf("Retrieved DoctorID from session: %d", doctorID)

	var doctorCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorID = ?", doctorID).Scan(&doctorCount); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error querying doctor count: %v", err)
		return
	}
	if doctorCount == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Doctor with provided DoctorID does not exist"})
		return
	}

	// Parse the frontend's datetime string into a time.Time object
	slotTime, err := time.Parse(time.RFC3339, slotDateTime)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format"})
		log.Printf("Error parsing datetime: %v", err)
		return
	}

	// Format the time.Time object to the required MySQL datetime format
	formattedSlotDateTime := slotTime.Format("2006-01-02 15:04:05")

	_, err = db.Exec("INSERT INTO slot (SlotDateTime, DoctorID) VALUES (?, ?)", formattedSlotDateTime, doctorID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Printf("Error inserting data into the database: %v", err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Slot added successfully!"
	log.Print("Slot added to the database")

	c.JSON(http.StatusOK, response)
}
