package health

type HealthUseCase interface {
	Readiness() error
}
