package main

import (
	"Clinic-Reservation-System/config"
	"Clinic-Reservation-System/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.DatabaseConnection()
	r := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200",
	"https://clinic-reservation-system-frnt-git-omarwael20021-dev.apps.sandbox-m3.1530.p1.openshiftapps.com"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	r.Use(cors.New(corsConfig))

	// Routes
	r.POST("/api/DoctorSignUp", controller.DoctorSignUp)
	r.POST("/api/DoctorSignIn", controller.DoctorSignIn)
	r.POST("/api/PatientSignUp", controller.PatientSignUp)
	r.POST("/api/PatientSignIn", controller.PatientSignIn)
	r.POST("/api/DoctorSetSchedule/:id", controller.SetSchedule)
	r.GET("/api/PatientShowAllDoctors", controller.ShowAllDoctors)
	r.GET("/api/PatientShowDoctorSlots", controller.ShowDoctorSlots)
	r.PUT("/api/PatientReserveSlot/:id", controller.ReserveSlot)
	r.PUT("/api/PatientUpdateAppointment/:id", controller.UpdateAppointment) 
	r.DELETE("/api/PatientCancelAppointment/:id", controller.CancelAppointment)
	r.GET("/api/PatientShowAppointments/:id", controller.ShowAllReservations) 

	port := ":8080"
	fmt.Println("Connected to port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
