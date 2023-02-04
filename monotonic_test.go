package monotonic

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	const maxIterations = 100_000_000
	t.Run("using monotonic time can not cause time shift", func(t *testing.T) {
		prev := Now()
		for i := 0; i < maxIterations; i++ {
			time.Sleep(time.Nanosecond)
			now := Now()
			// do NOT use Sub fn to verify where Sub is using monotonic time too
			if prev.UnixNano() >= now.UnixNano() {
				t.Fatalf(
					"only wall time(%d) might be greater than a time in the future(%d)",
					prev.UnixNano(),
					now.UnixNano(),
				)
			}
			prev = now
		}
	})

	t.Run("using wall time might cause time shift", func(t *testing.T) {
		prooved := false
		prev := time.Now()
		for i := 0; i < maxIterations; i++ {
			time.Sleep(1 * time.Nanosecond)
			now := time.Now()
			// do NOT use Sub fn to verify where Sub is using monotonic time too
			if prev.UnixNano() >= now.UnixNano() {
				prooved = true
				t.Logf(
					"previous wall time(%d) was greater than time.Now(%d) at iteration %d",
					prev.UnixNano(),
					now.UnixNano(),
					i,
				)
				break
			}
			prev = now
		}

		if !prooved {
			t.Errorf("could not prove!")
		}
	})
}
