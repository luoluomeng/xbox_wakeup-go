package main

import (
	"fmt"
	"testing"
)

func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestGenerateWakePackage(t *testing.T) {
	correctResult := []byte{0xdd, 0x02, 0x00, 0x13, 0x00, 0x00, 0x00, 0x10, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x00}
	liveid := "0123456789ABCDEF"
	wakepackage := generateWakePackage(liveid)
	if !ByteSliceEqual(correctResult, wakepackage) {
		t.Error("package generated is not match, test fail")
		fmt.Printf("package generated: %#x \n", wakepackage)
		fmt.Printf("correct package  : %#x \n", correctResult)
	}
}
