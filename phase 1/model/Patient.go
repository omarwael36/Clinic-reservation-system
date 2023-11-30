package model

type Patient struct {
	ID       int    `form:"patientId" json:"patientId"`
	Name     string `form:"userName" json:"userName"`
	Email    string `form:"userEmail" json:"userEmail"`
	Password string `form:"userPassword" json:"userPassword"`
}