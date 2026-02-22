package assert

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/internal"
)

// AlmostEqualTime is a function that compares two time values, time1 and time2,
// with a specified precision. This function is Soft Assert.
func AlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) bool {
	isAlmostEqual := internal.AlmostEqualTime(time1, time2, precision)
	if !isAlmostEqual {
		t.Error(internal.ShowDiffAlmostEqualTime(time1, time2, precision))
		return false
	}
	return true
}
