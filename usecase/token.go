package usecase

import (
	"context"
	"time"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/internal"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// RefreshTokenUsecase ...
func RefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return internal.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return internal.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return internal.ExtractIDFromToken(requestToken, secret)
}
