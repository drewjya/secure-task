package auth_service

import (
	domain "github.com/drewjya/secure-task/domain/auth"
	"github.com/drewjya/secure-task/repository"
)

type AuthService interface {
	LoginService(loginParam *domain.LoginDTO) (*string, error)
	RegisterService(registerParam *domain.RegisterDTO) (*string, error)
}

type AuthServiceImpl struct {
	queries *repository.Queries
}

func NewAuthService(
	queries *repository.Queries,
) AuthService {
	return &AuthServiceImpl{
		queries: queries,
	}
}
