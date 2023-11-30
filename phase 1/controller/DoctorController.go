package controller

import (
	"Clinic-Reservation-System/config"
	"Clinic-Reservation-System/helper"
	"Clinic-Reservation-System/model"
	"log"
	"net/http"

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
	helper.StoreUserInSession(c.Writer, c.Request, DoctorID, "Doctor")
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

	doctorID, userType := helper.GetUserFromSession(c.Request)
	if userType != "Doctor" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	slotDateTime, err := input.ParseSlotDateTime()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		log.Printf("Error parsing slotDateTime: %v", err)
		return
	}

	log.Printf("Retrieved DoctorID from session: %d", doctorID)
	var doctorCount int
	db.QueryRow("SELECT COUNT(*) FROM doctor WHERE DoctorID = ?", doctorID).Scan(&doctorCount)
	if doctorCount == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Doctor with provided DoctorID does not exist"})
		return
	}

	_, err = db.Exec("INSERT INTO slot (SlotDateTime, DoctorID) VALUES (?, ?)", slotDateTime, doctorID)
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
