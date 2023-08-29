package domain

import (
	"github.com/drewjya/secure-task/repository"
	"github.com/google/uuid"
)

type CreateAccountDTO struct {
	UserID   uuid.UUID `json:"user_id" validate:"required" `
	UserName string    `json:"user_name" validate:"required"`
}

type UpdateAccountDTO struct {
	UserID   uuid.UUID `json:"user_id" validate:"required" `
	UserName string    `json:"user_name" validate:"required"`
	Picture  *string   `json:"picture" `
}

func (account *CreateAccountDTO) MapToCreateAccount() repository.CreateAccountParams {
	return repository.CreateAccountParams{
		UserID:   account.UserID,
		UserName: account.UserName,
	}
}
func (account *UpdateAccountDTO) MapToUpdateAccount() repository.UpdateAccountParams {
	return repository.UpdateAccountParams{
		// Picture: ,
		ID:      account.UserID,
		Picture: account.Picture,

		UserName: account.UserName,
	}
}
