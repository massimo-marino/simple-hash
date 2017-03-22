package util

import (
	//	"math/rand"
	"testing"
	//	"time"
)

func TestSimpleHash_01(t *testing.T) {
	r := uint64(37)
	p := uint64((1 << 61) - 1)
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	hv1 := Hash(s, p, r)
	t.Logf("%s -> %v", s, hv1)

	hv2 := Hash(s, p, r)
	t.Logf("%s -> %v", s, hv2)

	if hv1 != hv2 {
		t.Fatal("Hash values for the same input key are different: They must be equal")
	}
}

func TestSimpleHashClosure_02(t *testing.T) {
	// generate the hash closure
	p := uint64((1 << 61) - 1)
	hf := HashClosure(p)

	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	hv1 := hf(s)
	t.Logf("%s -> %v", s, hv1)

	hv2 := hf(s)
	t.Logf("%s -> %v", s, hv2)

	if hv1 != hv2 {
		t.Fatal("Hash values for the same input key are different: They must be equal")
	}
}
