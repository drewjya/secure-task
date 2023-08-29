package account_service

import (
	"github.com/drewjya/secure-task/repository"
)

type AccountService interface {
}

type AccountServiceImpl struct {
	queries *repository.Queries
}

func NewAuthService(
	queries *repository.Queries,
) AccountService {
	return &AccountServiceImpl{
		queries: queries,
	}
}
