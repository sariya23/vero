package check

import (
	"sync"
	"testing"
	"time"
)

func TestAssertAlmostEqualTime(t *testing.T) {
	t.Parallel()
	timeNow := time.Now()
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Fatal(err)
	}
	timeNotUTC := timeNow.In(loc)
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
			name:      "time1 < time2 and those is equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Hour * 2,
			expected:  true,
		},
		{
			name:      "time1 < time2 and those is not equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Minute * 30,
			expected:  false,
		},
		{
			name:      "time1 > time2 and those is equal",
			time1:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			precision: time.Hour * 2,
			expected:  true,
		},
		{
			name:      "time1 > time2 and those is not equal",
			time1:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			precision: time.Minute * 30,
			expected:  false,
		},
		{
			name:      "the function ignores location information and calculates the difference in UTC",
			time1:     time.Now().UTC(),
			time2:     timeNotUTC,
			precision: time.Second * 20,
			expected:  false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockT := new(testing.T)
			res := AssertAlmostEqualTime(mockT, tc.time1, tc.time2, tc.precision)
			if res != tc.expected {
				t.Errorf("got %v, want %v", res, tc.expected)
			}
		})
	}
}

type mockTB struct {
	*testing.T

	mu              sync.Mutex
	failedNowCalled bool
}

func (m *mockTB) FailNow() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.failedNowCalled = true
}

func (*mockTB) Error(_ ...interface{}) {

}

func (m *mockTB) FailedNowCalled() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.failedNowCalled
}

func TestRequireAlmostEqualTime(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name         string
		time1, time2 time.Time
		precision    time.Duration
		isFailed     bool
	}{
		{
			name:      "time is almost equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Hour * 2,
			isFailed:  false,
		},
		{
			name:      "time is not almost equal",
			time1:     time.Date(2020, time.April, 15, 12, 0, 0, 0, time.UTC),
			time2:     time.Date(2020, time.April, 15, 13, 0, 0, 0, time.UTC),
			precision: time.Second * 2,
			isFailed:  true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockT := &mockTB{T: t}

			done := make(chan struct{})
			go func() {
				defer close(done)
				RequireAlmostEqualTime(mockT, tc.time1, tc.time2, tc.precision)
			}()
			<-done
			if mockT.FailedNowCalled() && !tc.isFailed {
				t.Errorf("FailNow was called but isFailed was false")
			}
			if !mockT.FailedNowCalled() && tc.isFailed {
				t.Errorf("FailNow was not called but isFailed was true")
			}
		})
	}
}
