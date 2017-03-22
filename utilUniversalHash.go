package util

import (
	"fmt"
	"math/rand"
	"time"
)

// Universal Hashing
//
// See the following links:
//
// https://www.cs.princeton.edu/courses/archive/fall09/cos521/Handouts/universalclasses.pdf
// https://en.wikipedia.org/wiki/Universal_hashing
//
// Simple implementation of universal hashing
// P >> M
// A, B are chosen randomly in the following intervals:
// 1 <= A <= P - 1
// 0 <= B <= P - 1
func simpleUniversalHash(key uint64, A uint64, B uint64, P uint64, M uint64) uint64 {
	return ((((A) * key) + B) % P) % M
}

// Simple implementation of universal hashing
// Implementation of Carter-Wegman method
//
// Possible DIGS values: 31, 61, 89
const DIGS = uint64(61)

// MARSENNEP = 2^DIGS - 1
const MARSENNEP = uint64((1 << DIGS) - 1)

// Choose MARSENNEP and M such that MARSENNEP >> M
// A, B are chosen randomly in the following intervals:
// 1 <= A <= MARSENNEP - 1
// 0 <= B <= MARSENNEP - 1
func universalHash(key uint64, A uint64, B uint64, M uint64) uint64 {
	hashVal := uint64((A)*key + B)

	hashVal = ((hashVal >> DIGS) + (hashVal & MARSENNEP))

	if hashVal >= MARSENNEP {
		hashVal = hashVal - MARSENNEP
	}

	return uint64(hashVal % M)
}

// A universal hash closure implementing the Carter-Wegman method
//
// Choose digs so that marsennep = 2^digs - 1 is a Marsenne prime number,
// and M such that marsennep >> M
// A, B are chosen randomly in the following intervals:
// 1 <= A <= marsennep - 1
// 0 <= B <= marsennep - 1
func UniversalHashClosure(digs uint64, M uint64) func(uint64) uint64 {
	marsennep := uint64((1 << digs) - 1)
	rand.Seed(time.Now().UnixNano())
	A := uint64(1 + rand.Int63n(int64(marsennep-1)))
	B := uint64(rand.Int63n(int64(marsennep)))

	fmt.Println("marsennep:", marsennep, "A:", A, " B:", B)

	return func(key uint64) uint64 {
		hashVal := uint64((A)*key + B)

		hashVal = ((hashVal >> digs) + (hashVal & marsennep))

		if hashVal >= marsennep {
			hashVal = hashVal - marsennep
		}

		return uint64(hashVal % M)
	}
}
