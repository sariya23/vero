package check

import (
	"testing"
	"time"

	"github.com/sariya23/vero/internal"
)

// AssertAlmostEqualTime is a function that compares two time values, time1 and time2,
// with a specified precision. This function is Soft Assert.
//
// The comparison ignores location information.
func AssertAlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) bool {
	time1 = time.Date(time1.Year(), time1.Month(), time1.Day(),
		time1.Hour(), time1.Minute(), time1.Second(), time1.Nanosecond(), time.UTC)
	time2 = time.Date(time2.Year(), time2.Month(), time2.Day(),
		time2.Hour(), time2.Minute(), time2.Second(), time2.Nanosecond(), time.UTC)
	var isAlmostEqual bool
	if time1.After(time2) {
		isAlmostEqual = time1.Sub(time2) < precision
	} else {
		isAlmostEqual = time2.Sub(time1) < precision
	}
	if !isAlmostEqual {
		t.Error(internal.ShowDiffAlmostEqualTime(time1, time2, precision))
		return false
	}
	return true
}

// RequireAlmostEqualTime compares two time values, time1 and time2,
// using the specified precision. If the difference exceeds the
// precision, the function fails the test immediately.
//
// The comparison ignores location information.
func RequireAlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) {
	isAlmostEqual := AssertAlmostEqualTime(t, time1, time2, precision)
	if !isAlmostEqual {
		t.FailNow()
	}
}
