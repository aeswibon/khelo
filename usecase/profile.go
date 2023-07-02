package usecase

import (
	"context"
	"time"

	"github.com/cp-Coder/khelo/domain"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// ProfileUsecase ...
func ProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Gender:   user.Gender,
		Age:      user.Age,
	}, nil
}

func (pu *profileUsecase) UpdateProfile(c context.Context, userID string, profile *domain.User) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	_, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	pu.userRepository.Update(ctx, userID, profile)
	return nil
}
