package collector

import "context"

type UseCase interface {
	UpdateStats(ctx context.Context, log *CollectorLog, ua string) error
}
