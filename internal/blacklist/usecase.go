package blacklist

import (
	"context"
)

type UseCase interface {
	IsIPBlacklisted(ctx context.Context, ip uint32) (bool, error)
	IsUABlacklisted(ctx context.Context, ui string) (bool, error)
}
