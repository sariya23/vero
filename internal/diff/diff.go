package diff

import (
	"fmt"
	"time"
)

// ShowDiffAlmostEqualTime formats time1, time2, and the precision for diff output.
func ShowDiffAlmostEqualTime(time1, time2 time.Time, precision time.Duration) string {
	showBuilder := "\ntime%d: year=%v, month=%v, day=%v, hour=%v, min=%v, second=%v, nanosecond=%v, tzinfo=(name=%v offset=%v)"
	time1Zone, time1Offset := time1.Zone()
	time2Zone, time2Offset := time2.Zone()
	time1Str := fmt.Sprintf(showBuilder, 1, time1.Year(), time1.Month(), time1.Day(), time1.Hour(), time1.Minute(), time1.Second(), time1.Nanosecond(), time1Zone, time1Offset)
	time2Str := fmt.Sprintf(showBuilder, 2, time2.Year(), time2.Month(), time2.Day(), time2.Hour(), time2.Minute(), time2.Second(), time2.Nanosecond(), time2Zone, time2Offset)
	return fmt.Sprintf("Time values is not equal with specified precision:%s%s\nprecision: %v", time1Str, time2Str, precision)
}
