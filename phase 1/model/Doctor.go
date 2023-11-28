package model

type Doctor struct {
	Id       int    `form:"DoctorId" json:"DoctorId"`
	Name     string `form:"userName" json:"userName"`
	Email    string `form:"userEmail" json:"userEmail"`
	Password string `form:"userPassword" json:"userPassword"`
}
