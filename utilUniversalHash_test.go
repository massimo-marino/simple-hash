package util

import (
	"math/rand"
	"testing"
	"time"
)

func TestUniversalHash_01(t *testing.T) {
	// M = 2^50
	M := uint64(1 << 50)
	// UniversalHashClosure() will use the Marsenne prime number 2^digs - 1
	digs := uint64(61)
	// Generate a universal hash closure
	uhf := UniversalHashClosure(digs, M)

	// limit should be set to M - 1, but the test would take hours, if not days
	limit := uint64(100000)
	rand.Seed(time.Now().UnixNano())
	for c := uint64(1); c <= limit; c++ {
		// Generate a random key in 0 <= key <= M-1
		key := uint64(rand.Int63n(int64(M)))

		// Apply the universal hash closure to key
		hv1 := uhf(key)
		t.Logf("%v -> %v", key, hv1)

		// Apply again the universal hash closure to the same key
		hv2 := uhf(key)
		t.Logf("%v -> %v", key, hv2)

		if hv1 != hv2 {
			t.Fatal("Hash values for the same input key are different: They must be equal")
		}
	}
}

func TestUniversalHash_02(t *testing.T) {
	// M = 2^50
	M := uint64(1 << 50)
	// UniversalHashClosure() will use the Marsenne prime number 2^digs - 1
	digs := uint64(61)
	// Generate a universal hash closure
	uhf := UniversalHashClosure(digs, M)

	// Apply the universal hash closure to key 0
	hv1 := uhf(uint64(0))
	// limit should be set to M - 1, but the test would take hours, if not days
	limit := uint64(1000000)
	for key := uint64(1); key <= limit; key++ {
		// Apply the universal hash closure to key
		hv2 := uhf(key)

		if hv1 == hv2 {
			t.Logf("%v -> %v", uint64(0), hv1)
			t.Logf("%v -> %v", key, hv2)
			t.Fatal("Hash values for two different input keys are equal: They must be different")
		}
	}
}

// Universal hash functions exist for strings also.
// 1. Choose any prime p, larger than M (and larger than the largest
// character code).

// 2. Use simple string hashing function, choosing the multiplier randomly
// between 1 and p − 1 and returning an intermediate hash value between
// 0 and p − 1, inclusive.

// 3. Apply a universal hash function to generate the ﬁnal hash value
// between 0 and M − 1
func TestUniversalHash_ForStrings_01(t *testing.T) {
	// 1. Choose any prime p, larger than M (and larger than the largest
	// character code).
	// M = 2^50
	M := uint64(1 << 50)
	// UniversalHashClosure() will use the Marsenne prime number 2^digs - 1
	digs := uint64(61)
	// p is the Marsenne prime 2^digs - 1: p >> M
	p := uint64((1 << digs) - 1)

	// 2. Use simple string hashing function, choosing the multiplier randomly
	// between 1 and p − 1 and returning an intermediate hash value between
	// 0 and p − 1, inclusive.
	// Generate a hash closure: the multiplier is chosen randomly in the
	// interval 1 <= multiplier <= p - 1
	hf := HashClosure(p)

	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	hv1 := hf(key)
	t.Logf("%s -> %v", key, hv1)

	if !((hv1 >= 0) && (hv1 <= (p - 1))) {
		t.Fatal("The intermediate hashed value is not in the right range")
	}

	// 3. Apply a universal hash function to generate the ﬁnal hash value
	// between 0 and M − 1
	// Generate a universal hash closure
	uhf := UniversalHashClosure(digs, M)

	// Apply the universal hash function to the intermediate hash value
	hv2 := uhf(hv1)
	t.Logf("%v -> %v", hv1, hv2)

	if !((hv2 >= 0) && (hv2 <= (M - 1))) {
		t.Fatal("The final hashed value is not in the right range")
	}

	// Apply again the universal hash function to the same intermediate hash value
	hv3 := uhf(hv1)
	t.Logf("%v -> %v", hv1, hv3)

	if hv2 != hv3 {
		t.Fatal("Final hash values for the same input key are different: They must be equal")
	}
}
