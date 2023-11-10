package model

import "time"

type Slot struct {
	SlotID       int    `json:"slotId"`
	SlotDateTime string `json:"slotDateTime"`
	DoctorID     int    `json:"doctorId"`
	PatientID    int    `json:"patientId"`
}

// Add a method to convert the string to time.Time
func (s *Slot) ParseSlotDateTime() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s.SlotDateTime)
}
