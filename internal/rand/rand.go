package rand

import (
	"fmt"
	mrand "math/rand/v2"
	"time"
)

// Source is a source of pseudo-random numbers (wrapper around math/rand/v2.Rand).
type Source struct {
	r *mrand.Rand
}

// NewRandSource creates a new Source instance seeded from the current time.
func NewRandSource() *Source {
	r := mrand.New(mrand.NewPCG(uint64(time.Now().UnixNano()),
		uint64(time.Now().UnixNano()>>32)))
	return &Source{r: r}
}

// IntRange returns a random int in [from, to] inclusive. Panics if from > to.
func (s *Source) IntRange(from, to int) int {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return s.r.IntN(to-from+1) + from
}

// Int8Range returns a random int8 in [from, to] inclusive. Panics if from > to.
func (s *Source) Int8Range(from, to int8) int8 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	n := int32(to) - int32(from) + 1
	return int8(s.r.Int32N(n) + int32(from))
}

// Int16Range returns a random int16 in [from, to] inclusive. Panics if from > to.
func (s *Source) Int16Range(from, to int16) int16 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	n := int32(to) - int32(from) + 1
	return int16(s.r.Int32N(n) + int32(from))
}

// Int32Range returns a random int32 in [from, to] inclusive. Panics if from > to.
func (s *Source) Int32Range(from, to int32) int32 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return s.r.Int32N(to-from+1) + from
}

// Int64Range returns a random int64 in [from, to] inclusive. Panics if from > to.
func (s *Source) Int64Range(from, to int64) int64 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return s.r.Int64N(to-from+1) + from
}

// UintRange returns a random uint in [from, to] inclusive. Panics if from > to.
func (s *Source) UintRange(from, to uint) uint {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return s.r.UintN(to-from+1) + from
}

// Uint8Range returns a random uint8 in [from, to] inclusive. Panics if from > to.
func (s *Source) Uint8Range(from, to uint8) uint8 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return uint8(s.r.Uint32N(uint32(to-from+1)) + uint32(from))
}

// Uint16Range returns a random uint16 in [from, to] inclusive. Panics if from > to.
func (s *Source) Uint16Range(from, to uint16) uint16 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return uint16(s.r.Uint32N(uint32(to-from+1)) + uint32(from))
}

// Uint32Range returns a random uint32 in [from, to] inclusive. Panics if from > to.
func (s *Source) Uint32Range(from, to uint32) uint32 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return s.r.Uint32N(to-from+1) + from
}

// Uint64Range returns a random uint64 in [from, to] inclusive. Panics if from > to.
func (s *Source) Uint64Range(from, to uint64) uint64 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return s.r.Uint64N(to-from+1) + from
}

// UintptrRange returns a random uintptr in [from, to] inclusive. Panics if from > to.
func (s *Source) UintptrRange(from, to uintptr) uintptr {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", from, to))
	}
	return uintptr(s.r.Uint64N(uint64(to-from+1))) + from
}

// Float32Range returns a random float32 in [from, to]. Panics if from > to.
func (s *Source) Float32Range(from, to float32) float32 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%f, %f]'", from, to))
	}
	return from + s.r.Float32()*(to-from)
}

// Float64Range returns a random float64 in [from, to]. Panics if from > to.
func (s *Source) Float64Range(from, to float64) float64 {
	if from > to {
		panic(fmt.Sprintf("invalid interval '[%f, %f]'", from, to))
	}
	return from + s.r.Float64()*(to-from)
}

// Int returns a random non-negative int.
func (s *Source) Int() int { return s.r.Int() }

// Int8 returns a random int8.
func (s *Source) Int8() int8 { return int8(s.r.Int32()) }

// Int16 returns a random int16.
func (s *Source) Int16() int16 { return int16(s.r.Int32()) }

// Int32 returns a random int32.
func (s *Source) Int32() int32 { return s.r.Int32() }

// Int64 returns a random int64.
func (s *Source) Int64() int64 { return s.r.Int64() }

// Uint returns a random uint.
func (s *Source) Uint() uint { return s.r.Uint() }

// Uint8 returns a random uint8.
func (s *Source) Uint8() uint8 { return uint8(s.r.Uint32()) }

// Uint16 returns a random uint16.
func (s *Source) Uint16() uint16 { return uint16(s.r.Uint32()) }

// Uint32 returns a random uint32.
func (s *Source) Uint32() uint32 { return s.r.Uint32() }

// Uint64 returns a random uint64.
func (s *Source) Uint64() uint64 { return s.r.Uint64() }

// Uintptr returns a random uintptr.
func (s *Source) Uintptr() uintptr { return uintptr(s.r.Uint64()) }

// Float32 returns a random float32 in [0, 1).
func (s *Source) Float32() float32 { return s.r.Float32() }

// Float64 returns a random float64 in [0, 1).
func (s *Source) Float64() float64 { return s.r.Float64() }

// Bool returns a random bool.
func (s *Source) Bool() bool {
	i := s.IntRange(0, 1)
	return i == 0
}
