package auth_service

import (
	"context"
	"log"

	domain "github.com/drewjya/secure-task/domain/auth"
	"github.com/drewjya/secure-task/libs/random"
	"github.com/drewjya/secure-task/repository"
)

func (re *AuthServiceImpl) RegisterService(registerParam *domain.RegisterDTO) (*string, error) {
	ctx := context.Background()

	userParams := registerParam.MapToCreateUser()
	log.Println(userParams)
	user, err := re.queries.CreateUser(ctx, userParams)
	if err != nil {
		return nil, err
	}
	session := random.GenerateRandomString(32)
	err = re.queries.CreateSession(ctx, repository.CreateSessionParams{
		UserID:       user.ID,
		SessionToken: session,
	})
	if err != nil {
		return nil, err
	}

	return &session, nil

}
