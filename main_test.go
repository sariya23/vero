package main

import (
	"testing"
	"time"

	"github.com/sariya23/probatigo/check/require"
)

func TestB(t *testing.T) {
	require.AlmostEqualTime(t, time.Now(), time.Now().UTC(), time.Second)
}
