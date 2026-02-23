package internal

import "time"

// AlmostEqualTime returns true if the difference between time1 and time2
// does not exceed the specified precision.
//
// This function ignore information about Location
func AlmostEqualTime(time1, time2 time.Time, precision time.Duration) bool {
	time1 = time.Date(time1.Year(), time1.Month(), time1.Day(),
		time1.Hour(), time1.Minute(), time1.Second(), time1.Nanosecond(), time.UTC)
	time2 = time.Date(time2.Year(), time2.Month(), time2.Day(),
		time2.Hour(), time2.Minute(), time2.Second(), time2.Nanosecond(), time.UTC)
	if time1.After(time2) {
		return time1.Sub(time2) < precision
	}
	return time2.Sub(time1) < precision
}
