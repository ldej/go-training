package main

import "testing"

func TestRaceCondition(t *testing.T) {
	// run this with go test -race to see:
	// ==================
	// WARNING: DATA RACE
	// Read at 0x000000726888 by goroutine 9:
	// --- FAIL: TestRaceCondition (0.00s)
	//    testing.go:1092: race detected during execution of test
	run()
}
