package util

import (
	"math/rand"
	"time"
)

// A hash routine for string objects
//
// p must be a prime number, and the multiplier must be:
// 1 <= multiplier <= p - 1
//
func Hash(key string, p uint64, multiplier uint64) uint64 {
	hashVal := uint64(0)

	for _, ch := range key {
		hashVal = multiplier*hashVal + uint64(ch)
	}
	return uint64(hashVal % p)

}

// A hash closure for string objects
//
// p must be a prime number, and the multiplier is chosen randomly such that:
// 1 <= multiplier <= p - 1
//
func HashClosure(p uint64) func(string) uint64 {
	rand.Seed(time.Now().UnixNano())
	multiplier := uint64(1 + rand.Int63n(int64(p-1)))
	return func(key string) uint64 {
		hashVal := uint64(0)

		for _, ch := range key {
			hashVal = multiplier*hashVal + uint64(ch)
		}
		return uint64(hashVal % p)
	}
}
