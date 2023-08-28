package domain

import "github.com/drewjya/secure-task/repository"

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email" `
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterDTO struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (register *RegisterDTO) MapToCreateUser() repository.CreateUserParams {
	return repository.CreateUserParams{
		Email:    register.Email,
		Name:     register.Name,
		Password: register.Password,
	}
}
