package model

type Doctor struct {
	Id       int    `form:"DoctorId" json:"DoctorId"`
	Name     string `form:"DoctorName" json:"DoctorName"`
	Email    string `form:"DoctorEmail" json:"DoctorEmail"`
	Password string `form:"DoctorPassword" json:"DoctorPassword"`
}
