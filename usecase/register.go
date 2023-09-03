package usecase

import (
	"context"
	"log"
	"time"

	"github.com/cp-Coder/khelo/domain"
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

func (su *registerUsecase) Create(c context.Context, request *domain.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user := &domain.User{
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Age:      request.Age,
		Password: request.Password,
	}

	err := su.userRepository.Create(ctx, user)
	log.Default().Println(err)
	return err
}
