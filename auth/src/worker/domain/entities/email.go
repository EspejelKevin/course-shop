package entities

type Email struct {
	Name    string `json:"name" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	From    string `json:"from" binding:"required,email"`
	To      string `json:"to" binding:"required,email"`
	Code    string `json:"code" binding:"required,len=20"`
}
