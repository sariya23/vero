package rand

import (
	"fmt"
	mrand "math/rand/v2"
	"time"
)

type RandSource struct {
	r *mrand.Rand
}

func NewRandSource() *RandSource {
	r := mrand.New(mrand.NewPCG(uint64(time.Now().UnixNano()),
		uint64(time.Now().UnixNano()>>32)))
	return &RandSource{r: r}
}

func (s *RandSource) IntRange(min, max int) int {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return s.r.IntN(max-min+1) + min
}

func (s *RandSource) Int8Range(min, max int8) int8 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	n := int32(max) - int32(min) + 1
	return int8(s.r.Int32N(n) + int32(min))
}

func (s *RandSource) Int16Range(min, max int16) int16 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	n := int32(max) - int32(min) + 1
	return int16(s.r.Int32N(n) + int32(min))
}

func (s *RandSource) Int32Range(min, max int32) int32 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return s.r.Int32N(max-min+1) + min
}

func (s *RandSource) Int64Range(min, max int64) int64 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return s.r.Int64N(max-min+1) + min
}

func (s *RandSource) UintRange(min, max uint) uint {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return s.r.UintN(max-min+1) + min
}

func (s *RandSource) Uint8Range(min, max uint8) uint8 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return uint8(s.r.Uint32N(uint32(max-min+1)) + uint32(min))
}

func (s *RandSource) Uint16Range(min, max uint16) uint16 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return uint16(s.r.Uint32N(uint32(max-min+1)) + uint32(min))
}

func (s *RandSource) Uint32Range(min, max uint32) uint32 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return s.r.Uint32N(max-min+1) + min
}

func (s *RandSource) Uint64Range(min, max uint64) uint64 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return s.r.Uint64N(max-min+1) + min
}

func (s *RandSource) UintptrRange(min, max uintptr) uintptr {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%d, %d]'", min, max))
	}
	return uintptr(s.r.Uint64N(uint64(max-min+1))) + min
}

func (s *RandSource) Float32Range(min, max float32) float32 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%f, %f]'", min, max))
	}
	return min + s.r.Float32()*(max-min)
}

func (s *RandSource) Float64Range(min, max float64) float64 {
	if min > max {
		panic(fmt.Sprintf("invalid interval '[%f, %f]'", min, max))
	}
	return min + s.r.Float64()*(max-min)
}

func (s *RandSource) Int() int     { return s.r.Int() }
func (s *RandSource) Int8() int8   { return int8(s.r.Int32()) }
func (s *RandSource) Int16() int16 { return int16(s.r.Int32()) }
func (s *RandSource) Int32() int32 { return s.r.Int32() }
func (s *RandSource) Int64() int64 { return s.r.Int64() }

func (s *RandSource) Uint() uint       { return s.r.Uint() }
func (s *RandSource) Uint8() uint8     { return uint8(s.r.Uint32()) }
func (s *RandSource) Uint16() uint16   { return uint16(s.r.Uint32()) }
func (s *RandSource) Uint32() uint32   { return s.r.Uint32() }
func (s *RandSource) Uint64() uint64   { return s.r.Uint64() }
func (s *RandSource) Uintptr() uintptr { return uintptr(s.r.Uint64()) }

func (s *RandSource) Float32() float32 { return s.r.Float32() }
func (s *RandSource) Float64() float64 { return s.r.Float64() }

func (s *RandSource) Bool() bool {
	i := s.IntRange(0, 1)
	return i == 0
}
