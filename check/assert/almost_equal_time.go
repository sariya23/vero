package assert

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/internal"
)

// AlmostEqualTime is a function that compares two time values, time1 and time2,
// with a specified precision. This function is Soft Assert.
func AlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) {
	isAlmostEqual := internal.AlmostEqualTime(time1, time2, precision)
	if isAlmostEqual {
		t.Errorf("time1 is '%v' and time2 is '%v' is not equal with precision '%v'",
			time1, time2, precision)
	}
}
