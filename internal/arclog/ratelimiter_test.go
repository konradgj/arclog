package arclog

import (
	"testing"
	"time"
)

func TestRateLimiter_AllowsWithinLimit(t *testing.T) {
	rl := NewRateLimiter(2, time.Second)
	rl.Wait()
	rl.Wait()

	done := make(chan struct{})
	go func() {
		rl.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatalf("expected third Allow to block, but it returned immediately")
	case <-time.After(100 * time.Millisecond):
	}
}

func TestRateLimiter_Refills(t *testing.T) {
	rl := NewRateLimiter(1, 200*time.Millisecond)
	rl.Wait()

	start := time.Now()
	rl.Wait()

	elapsed := time.Since(start)
	if elapsed < 150*time.Millisecond {
		t.Errorf("expected Allow to block ~200ms, but returned after %v", elapsed)
	}
}
