package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestShowDiffAlmostEqualTime(t *testing.T) {
	t.Parallel()
	time1 := time.Now()
	time2 := time.Now().UTC()
	precision := time.Second
	expectedTzNameTime1, expectedTzOffsetTime1 := time1.Zone()
	expectedTime1 := fmt.Sprintf("\ntime1: year=%v, month=%v, day=%v, hour=%v, min=%v, second=%v, nanosecond=%v, tzinfo=(name=%v offset=%v)",
		time1.Year(), time1.Month(), time1.Day(), time1.Hour(), time1.Minute(), time1.Second(), time1.Nanosecond(), expectedTzNameTime1, expectedTzOffsetTime1)
	expectedTzNameTime2, expectedTzOffsetTime2 := time2.Zone()
	expectedTime2 := fmt.Sprintf("\ntime2: year=%v, month=%v, day=%v, hour=%v, min=%v, second=%v, nanosecond=%v, tzinfo=(name=%v offset=%v)",
		time2.Year(), time2.Month(), time2.Day(), time2.Hour(), time2.Minute(), time2.Second(), time2.Nanosecond(), expectedTzNameTime2, expectedTzOffsetTime2)
	expected := fmt.Sprintf("Time values is not equal with specified precision:%s%s\nprecision: %v", expectedTime1, expectedTime2, precision)

	got := ShowDiffAlmostEqualTime(time1, time2, precision)
	if expected != got {
		t.Errorf("got\n'%s' expected\n'%s'", got, expected)
	}
}
