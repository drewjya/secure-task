package auth_service

import (
	"context"
	"errors"
	"log"

	domain "github.com/drewjya/secure-task/domain/auth"
	"github.com/drewjya/secure-task/libs/hash"
	"github.com/drewjya/secure-task/libs/random"
	"github.com/drewjya/secure-task/repository"
)

func (re *AuthServiceImpl) LoginService(loginParam *domain.LoginDTO) (*string, error) {
	ctx := context.Background()
	log.Println("asjsajsj")
	user, err := re.queries.GetOneUserWithEmail(ctx, loginParam.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	checkSession, err := re.queries.GetLatestSessionWithUserId(ctx, user.ID)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	re.queries.InvalidateTokenLatest(ctx, checkSession.ID)
	verifyHash := hash.CheckPasswordHash(loginParam.Password, user.Password)
	session := random.GenerateRandomString(32)
	err = re.queries.CreateSession(ctx, repository.CreateSessionParams{
		UserID:       user.ID,
		SessionToken: session,
	})
	if err != nil {
		return nil, err
	}

	if verifyHash {
		return &session, nil

	}
	return nil, errors.New("wrong password")
}
