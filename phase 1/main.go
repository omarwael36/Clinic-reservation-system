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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"} // Update with your Angular app's URL
	r.Use(cors.New(config))

	r.POST("/api/DoctorSignUp", controller.DoctorSignUp)
	r.GET("/api/DoctorSignIn", controller.DoctorSignIn)
	r.POST("/api/PatientSignUp", controller.PatientSignUp)
	r.GET("/api/PatientSignIn", controller.PatientSignIn)
	r.POST("/api/DoctorSetSchedule", controller.SetSchedule)
	r.GET("/api/PatientShowAllDoctors", controller.ShowAllDoctors)
	r.GET("/api/PatientShowDoctorSlots", controller.ShowDoctorSlots)
	r.PUT("/api/PatientReserveSlot", controller.ReserveSlot)
	r.PUT("/api/PatientUpdateAppointment", controller.UpdateAppointment)
	r.DELETE("/api/PatientCancelAppointment", controller.CancelAppointment)
	r.GET("/api/PatientShowAppointments", controller.ShowAllReservations)

	fmt.Println("connected to port 8080")
	r.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
