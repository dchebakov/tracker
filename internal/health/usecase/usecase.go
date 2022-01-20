package usecase

import "github.com/dchebakov/tracker/internal/health"

type healthUC struct {
	healthRepo health.Repository
}

func NewHealthUseCase(health health.Repository) *healthUC {
	return &healthUC{healthRepo: health}
}

func (u *healthUC) Readiness() error {
	return u.healthRepo.Readiness()
}
