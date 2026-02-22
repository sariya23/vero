package internal

import "time"

// AlmostEqualTime returns true if the difference between time1 and time2
// does not exceed the specified precision.
func AlmostEqualTime(time1, time2 time.Time, precision time.Duration) bool {
	if time1.After(time2) {
		return time1.Sub(time2) < precision
	}
	return time2.Sub(time1) < precision
}
