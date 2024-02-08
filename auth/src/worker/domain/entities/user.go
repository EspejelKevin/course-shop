package entities

type User struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Celular  string `json:"celular" validate:"required"`
	Rol      string `json:"rol" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	Verified bool   `json:"verified" default:"false"`
}
