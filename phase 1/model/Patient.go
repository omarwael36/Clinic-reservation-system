package model

type Patient struct {
	ID       int    `json:"patientId"`
	Name     string `json:"patientName"`
	Email    string `json:"patientEmail"`
	Password string `json:"patientPassword"`
}
