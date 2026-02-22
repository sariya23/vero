package internal

import (
	"testing"
	"time"
)

func TestAlmostEqualTime(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name         string
		time1, time2 time.Time
		precision    time.Duration
		expected     bool
	}{
		{
			name:      "time1 is equal to time2",
			time1:     time.Date(2020, time.April, 15, 0, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 0, 0, 0, 0, time.UTC),
			precision: time.Second,
			expected:  true,
		},
		{
			name:      "time1 less than time2 and those almost equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Hour * 2,
			expected:  true,
		},
		{
			name:      "time1 greater than time2 and those almost equal",
			time1:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			precision: time.Hour * 2,
			expected:  true,
		},
		{
			name:      "time1 and time2 is not almost equal",
			time1:     time.Date(2020, time.April, 15, 12, 20, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 12, 21, 0, 0, time.UTC),
			precision: time.Second * 20,
			expected:  false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := AlmostEqualTime(tc.time1, tc.time2, tc.precision)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
