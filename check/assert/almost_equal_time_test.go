package assert

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
			name:      "time is almost equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Hour * 2,
			expected:  true,
		},
		{
			name:      "time is not almost equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Second * 2,
			expected:  false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockT := new(testing.T)
			res := AlmostEqualTime(mockT, tc.time1, tc.time2, tc.precision)
			if res != tc.expected {
				t.Errorf("got %v, want %v", res, tc.expected)
			}
		})
	}
}
