package rands

import (
	"math/rand"
	"sync"
	"time"
)

var (
	sMu sync.Mutex
	src = rand.NewSource(time.Now().UnixNano())

	mu  sync.Mutex
	rad = rand.New(src)
)

// Int 随机int
func Int() (n int) {
	return rad.Int()
}

// SafeInt 线性安全生成一个随机整数
func SafeInt() (n int) {
	mu.Lock()
	n = Int()
	mu.Unlock()
	return
}

// IntBetween 在[min, max]之间随机一个整数
func IntBetween(min, max int) (n int) {
	return rad.Intn(max-min+1) + min
}

// SafeIntBetween 线性安全的生成随机数
func SafeIntBetween(min, max int) (n int) {
	mu.Lock()
	n = IntBetween(min, max)
	mu.Unlock()
	return
}

// Int32 随机生成int32
func Int32() (n int32) {
	return rad.Int31()
}

// SafeInt32 线性安全的随机生成int32
func SafeInt32() (n int32) {
	mu.Lock()
	n = Int32()
	mu.Unlock()
	return
}

// Int32Between 在[min, max]之间随机生成一个int32
func Int32Between(min, max int32) (n int32) {
	return rad.Int31n(max-min+1) + min
}

// SafeInt32Between 线性安全的Int32Between
func SafeInt32Between(min, max int32) (n int32) {
	mu.Lock()
	n = Int32Between(min, max)
	mu.Unlock()
	return
}

// Int64 随机生成int64
func Int64() (n int64) {
	return rad.Int63()
}

// SafeInt64 线性安全的随机生成int64
func SafeInt64() (n int64) {
	mu.Lock()
	n = Int64()
	mu.Unlock()
	return
}

// Int64Between 随机生成int64
func Int64Between(min, max int64) (n int64) {
	return rad.Int63n(max-min+1) + min
}

// SafeInt64Between 线性安全的随机生成int64
func SafeInt64Between(min, max int64) (n int64) {
	mu.Lock()
	n = Int64Between(min, max)
	mu.Unlock()
	return
}

// Float32 随机生成float32
func Float32() (f float32) {
	return rad.Float32()
}

// SafeFloat32 线性安全的随机生成float32
func SafeFloat32() (f float32) {
	mu.Lock()
	f = Float32()
	mu.Unlock()
	return
}

// Float64 随机生成float64
func Float64() (f float64) {
	return rad.Float64()
}

// SafeFloat64 线性安全的随机生成float64
func SafeFloat64() (f float64) {
	mu.Lock()
	f = Float64()
	mu.Unlock()
	return
}
