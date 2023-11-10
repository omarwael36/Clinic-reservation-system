package main

import (
	"Clinic-Reservation-System/config"
	"Clinic-Reservation-System/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Main function
func main() {
	config.DatabaseConnection()
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/DoctorSignUp", controller.DoctorSignUp).Methods("POST")
	router.HandleFunc("/DoctorSignIn", controller.DoctorSignIn).Methods("GET")
	router.HandleFunc("/PatientSignUp", controller.PatientSignUp).Methods("POST")
	router.HandleFunc("/PatientSignIn", controller.PatientSignIn).Methods("GET")
	router.HandleFunc("/DoctorSetSchedule", controller.SetSchedule).Methods("POST")
	router.HandleFunc("/PatientShowAllDoctors", controller.ShowAllDoctors).Methods("GET")
	router.HandleFunc("/PatientShowDoctorSlots", controller.ShowDoctorSlots).Methods("GET")
	router.HandleFunc("/PatientReserveSlot", controller.ReserveSlot).Methods("PUT")
	router.HandleFunc("/PatientUpdateAppointment", controller.UpdateAppointment).Methods("PUT")
	router.HandleFunc("/PatientCancelAppointment", controller.CancelAppointment).Methods("Delete")
	router.HandleFunc("/PatientShowAppointments", controller.ShowAllReservations).Methods("GET")
	fmt.Println("connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
