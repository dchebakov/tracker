package blacklist

import "context"

type Repository interface {
	HasIP(ctx context.Context, ip uint32) (bool, error)
	HasUA(ctx context.Context, ua string) (bool, error)
}
