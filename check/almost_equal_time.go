package check

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/internal"
)

// AssertAlmostEqualTime is a function that compares two time values, time1 and time2,
// with a specified precision. This function is Soft Assert.
//
// The comparison ignores location information.
func AssertAlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) bool {
	isAlmostEqual := internal.AlmostEqualTime(time1, time2, precision)
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
