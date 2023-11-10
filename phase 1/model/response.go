package model

type DoctorResponse struct {
	ID    int    `json:"DoctorId"`
	Name  string `json:"DoctorName"`
	Email string `json:"DoctorEmail"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
