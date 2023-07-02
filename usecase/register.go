package usecase

import (
	"context"
	"time"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/internal/tokenutil"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// RegisterUsecase ...
func RegisterUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *registerUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *registerUsecase) CheckUser(c context.Context, username string, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	if _, err := su.userRepository.GetUserByUsername(ctx, username); err != nil {
		return false, err
	}
	if _, err := su.userRepository.GetUserByEmail(ctx, email); err != nil {
		return false, err
	}
	return true, nil
}

func (su *registerUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *registerUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
