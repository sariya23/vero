package internal

import "time"

func AlmostEqualTime(time1, time2 time.Time, precision time.Duration) bool {
	if time1.After(time2) {
		return time1.Sub(time2) < precision
	}
	return time2.Sub(time1) < precision
}
