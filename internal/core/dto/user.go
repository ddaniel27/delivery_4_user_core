package dto

type CreateUserDTO struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Inst      string `json:"institution" binding:"required"`
	City      string `json:"city" binding:"required"`
	Birthdate string `json:"birthdate" binding:"required"`
}
