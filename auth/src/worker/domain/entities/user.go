package entities

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,alphanum"`
	Celular  string `json:"celular" binding:"required,e164"`
	Rol      string `json:"rol" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Verified bool   `json:"verified" default:"false"`
}
