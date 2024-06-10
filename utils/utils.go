package utils

import (
	"time"
	"math/rand"
)


func randomSeed() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// returns some float64 in the range [0..1)
func RandomFloat64() float64 {
	rng := randomSeed()
	return rng.Float64() 
}
	
// returns some float64 in the range [min..max)
func RandomFloat64Range(min, max float64) float64 {
	return min + RandomFloat64() * (max - min)
}

func Sleep(s int) {
	time.Sleep(time.Duration(s) * time.Second)
}