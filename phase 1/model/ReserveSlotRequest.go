package model

type ReserveSlotRequest struct {
	SlotID    int `json:"slotId"`
	PatientID int `json:"patientId"`
}
