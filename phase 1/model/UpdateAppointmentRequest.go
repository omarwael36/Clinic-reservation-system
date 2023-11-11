package model

type UpdateAppointmentRequest struct {
	AppointmentID int `json:"appointmentId"`
	NewSlotID     int `json:"newSlotId"`
	NewDoctorID   int `json:"newDoctorId"`
	NewPatientID  int `json:"newPatientId"`
}
