package monotonic

import "time"

var initializedAt time.Time

func init() {
	initializedAt = time.Now()
}

// Now returns monotonic clock time where the package initialization time is
// used as the reference
func Now() time.Time {
	d := time.Since(initializedAt)
	return initializedAt.Add(d)
}
