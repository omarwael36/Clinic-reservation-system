package model

import "time"

type Slot struct {
	SlotID       int    `form:"slotId" json:"slotId"`
	SlotDateTime string `form:"slotDateTime" json:"slotDateTime"`
	DoctorID     int    `form:"doctorId" json:"doctorId"`
	PatientID    int    `form:"patientId" json:"patientId"`
}

// Add a method to convert the string to time.Time
func (s *Slot) ParseSlotDateTime() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", s.SlotDateTime)
}
