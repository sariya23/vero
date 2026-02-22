package require

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/internal"
)

func AlmostEqualTime(t testing.TB, time1, time2 time.Time, precision time.Duration) {
	isAlmostEqual := internal.AlmostEqualTime(time1, time2, precision)
	if isAlmostEqual {
		t.Fatalf("time1 is '%v' and time2 is '%v' is not equal with precision '%v'",
			time1, time2, precision)
	}
}
