package util

import (
	"testing"
)

func TestSimpleHash_01(t *testing.T) {
	multiplier := uint64(37)
	// p is the prime 2^61 - 1; in general use a Marsenne prime number
	p := uint64((1 << 61) - 1)
	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	hv1 := Hash(key, p, multiplier)
	t.Logf("%s -> %v", key, hv1)

	hv2 := Hash(key, p, multiplier)
	t.Logf("%s -> %v", key, hv2)

	if hv1 != hv2 {
		t.Fatal("Hash values for the same input key are different: They must be equal")
	}
}

func TestSimpleHashClosure_01(t *testing.T) {
	// p is the prime 2^61 - 1; in general use a Marsenne prime number
	p := uint64((1 << 61) - 1)
	// generate the hash closure: the multiplier is chosen randomly in the
	// interval 1 <= multiplier <= p - 1
	hf := HashClosure(p)

	key := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	hv1 := hf(key)
	t.Logf("%s -> %v", key, hv1)

	hv2 := hf(key)
	t.Logf("%s -> %v", key, hv2)

	if hv1 != hv2 {
		t.Fatal("Hash values for the same input key are different: They must be equal")
	}
}
