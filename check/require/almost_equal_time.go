package require

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/check/assert"
)

// AlmostEqualTime is a function that compares two time values, time1 and time2,
// with a specified precision. This function is Fail Now.
func AlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) {
	isAlmostEqual := assert.AlmostEqualTime(t, time1, time2, precision)
	if !isAlmostEqual {
		t.FailNow()
	}
}
