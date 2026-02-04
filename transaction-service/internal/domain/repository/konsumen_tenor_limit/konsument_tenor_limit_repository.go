package konsumentenorlimit

import "context"

type KonsumenTenorLimitRepository interface {
	UpdateBalance(ctx context.Context, ) error
}