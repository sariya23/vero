package internal

import "time"

func AlmostEqualTime(a, b time.Time, precision time.Duration) bool {
	if a.After(b) {
		return a.Sub(b) < precision
	}
	return b.Sub(a) < precision
}
