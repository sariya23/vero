package require

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/check/assert"
)

// AlmostEqualTime compares two time values, time1 and time2,
// using the specified precision. If the difference exceeds the
// precision, the function fails the test immediately.
//
// The comparison ignores location information.
func AlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) {
	isAlmostEqual := assert.AlmostEqualTime(t, time1, time2, precision)
	if !isAlmostEqual {
		t.FailNow()
	}
}
