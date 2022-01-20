package usecase

import "github.com/dchebakov/tracker/internal/health"

type Health struct {
	healthRepo health.Repository
}

func NewHealthUseCase(health health.Repository) *Health {
	return &Health{healthRepo: health}
}

func (u *Health) Readiness() error {
	return u.healthRepo.Readiness()
}
