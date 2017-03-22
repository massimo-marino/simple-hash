package util

import (
	"math/rand"
	"time"
)

//
// This hash function involves all characters in the key and can generally be
// expected to distribute well.
// The code computes a polynomial function (of multiplier) by use of Horner’s rule
// (see https://en.wikipedia.org/wiki/Horner%27s_method) and brings the result
// into proper range.
// For instance, another way of computing hk = key[0] + multiplier * key[1] + multiplier^2 * k[2]
// is by the formula hk = ((k[2] ) ∗ multiplier + k[1] ) ∗ multiplier + k[0]
// Horner’s rule extends this to an n-th degree polynomial.
//
// The hash function takes advantage of the fact that overﬂow is allowed and
// uses unsigned int's to avoid introducing a negative number.
// The hash function implemented here is not necessarily the best with respect
// to table distribution, but it does have the merit of extreme simplicity and
// is reasonably fast.
// If the keys are very long, the hash function will take too long to compute.
// A common practice in this case is not to use all the characters. The length
// and properties of the keys would then inﬂuence the choice.
// For instance, the keys could be a complete street address. The hash function
// might include a couple of characters from the street address and perhaps a
// couple of characters from the city name and ZIP code.
// Some programmers implement their hash function by using only the characters
// in the odd spaces, with the idea that the time saved computing the hash
// function will make up for a slightly less evenly distributed function.

// A hash routine for string objects
// p must be a prime number (a Marsenne prime number is the best choice), and
// the multiplier must be: 1 <= multiplier <= p - 1
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
// p must be a prime number (a Marsenne prime number is the best choice), and
// the multiplier is chosen randomly such that: 1 <= multiplier <= p - 1
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