package arclog

import "time"

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(limit int, per time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, limit),
	}

	go func() {
		ticker := time.NewTicker(per / time.Duration(limit))
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}()

	return rl
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}
