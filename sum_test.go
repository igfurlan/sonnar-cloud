package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestTimes(t *testing.T) {
	result := times(2, 3)

	if result != 6 {
		t.Error("The result must be 6")
	}
}

func TestSub(t *testing.T) {
	result := sub(3, 2)

	if result != 1 {
		t.Error("The result must be 1")
	}
}
func TestSumx(t *testing.T) {
	result := sumX(2, 3)

	if result != 7 {
		t.Error("The result must be 7")
	}
}
