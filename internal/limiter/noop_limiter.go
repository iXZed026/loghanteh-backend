package limiter

import "context"

type NoopLimiter struct{}

func NewNoopLimiter() *NoopLimiter {
	return &NoopLimiter{}
}

func (n *NoopLimiter) Allow(
	_ context.Context,
	_ string,
	_ int,
) (bool, error) {
	return true, nil
}
