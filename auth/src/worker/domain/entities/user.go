package entities

type User struct {
	Id            int    `json:"-"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required,min=8,alphanum"`
	Phone         string `json:"phone" binding:"required,e164"`
	Rol           string `json:"rol" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Lastname      string `json:"lastname" binding:"required"`
	VerifiedEmail bool   `json:"-"`
}

type UserIdentity struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,alphanum"`
}

type VerificationCode struct {
	Code string `json:"code" binding:"required,len=20"`
}
