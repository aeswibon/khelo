package usecase

import (
	"context"
	"time"

	"github.com/cp-Coder/khelo/domain"
)

type facilityUsecase struct {
	facilityRepository domain.FacilityRepository
	contextTimeout     time.Duration
}

// FacilityUsecase ...
func FacilityUsecase(facilityRepository domain.FacilityRepository, timeout time.Duration) domain.FacilityUsecase {
	return &facilityUsecase{
		facilityRepository: facilityRepository,
		contextTimeout:     timeout,
	}
}

func (su *facilityUsecase) Create(c context.Context, facility *domain.Facility) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.facilityRepository.Create(ctx, facility)
}

func (su *facilityUsecase) Fetch(c context.Context) ([]domain.Facility, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.facilityRepository.Fetch(ctx)
}

func (su *facilityUsecase) GetFacilityByName(c context.Context, name string) (domain.Facility, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.facilityRepository.GetFacilityByName(ctx, name)
}

func (su *facilityUsecase) GetFacilityByEmail(c context.Context, email string) (domain.Facility, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.facilityRepository.GetFacilityByEmail(ctx, email)
}

func (su *facilityUsecase) GetFacilityByID(c context.Context, id string) (domain.Facility, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.facilityRepository.GetByID(ctx, id)
}

func (su *facilityUsecase) UpdateFacility(c context.Context, id string, facility *domain.Facility) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.facilityRepository.Update(ctx, id, facility)
}
