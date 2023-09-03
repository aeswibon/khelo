package usecase

import (
	"context"
	"time"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/internal"
	"go.mongodb.org/mongo-driver/bson"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// LoginUsecase ...
func LoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) Authenticate(c context.Context, request domain.LoginRequest) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	users, err := lu.userRepository.Fetch(ctx, bson.M{"username": request.Username}, bson.M{
		"username": 1,
	})
	if err != nil || len(users) == 0 {
		return domain.User{}, err
	}
	user := users[0]
	check := internal.CheckPasswordHash(request.Password, user.Password)
	if !check {
		return domain.User{}, err
	}
	return user, nil
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return internal.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return internal.CreateRefreshToken(user, secret, expiry)
}
