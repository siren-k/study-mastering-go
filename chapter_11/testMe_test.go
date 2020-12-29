package main

import "testing"

func TestS1(t *testing.T) {
	if s1("123456789") != 9 {
		t.Error(`s1("123456789") != 9`)
	}
	if s1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TestS2(t *testing.T) {
	if s2("123456789") != 9 {
		t.Error(`s2("123456789") != 9`)
	}
	if s2("") != 0 {
		t.Error(`s2("") != 0`)
	}
}

func TestF1(t *testing.T) {
	if f1(0) != 0 {
		t.Error(`f1(0) != 0`)
	}
	if f1(1) != 1 {
		t.Error(`f1(1) != 1`)
	}
	if f1(2) != 1 {
		t.Error(`f1(2) != 1`)
	}
	if f1(10) != 55 {
		t.Error(`f1(10) != 55`)
	}
}

func TestF2(t *testing.T) {
	if f2(0) != 0 {
		t.Error(`f2(0) != 0`)
	}
	if f2(1) != 1 {
		t.Error(`f2(1) != 1`)
	}
	if f2(2) != 1 {
		t.Error(`f2(2) != 1`)
	}
	if f2(10) != 55 {
		t.Error(`f2(10) != 55`)
	}
}

// ❯ go test testMe.go testMe_test.go -v
// === RUN   TestS1
//     testMe_test.go:7: s1("123456789") != 9
// --- FAIL: TestS1 (0.00s)
// === RUN   TestS2
// --- PASS: TestS2 (0.00s)
// === RUN   TestF1
// --- PASS: TestF1 (0.00s)
// === RUN   TestF2
//     testMe_test.go:43: f2(1) != 1
//     testMe_test.go:46: f2(2) != 1
//     testMe_test.go:49: f2(10) != 55
// --- FAIL: TestF2 (0.00s)
// FAIL
// FAIL    command-line-arguments  0.283s
// FAIL

// ❯ go test testMe.go testMe_test.go
// --- FAIL: TestS1 (0.00s)
//     testMe_test.go:7: s1("123456789") != 9
// --- FAIL: TestF2 (0.00s)
//     testMe_test.go:43: f2(1) != 1
//     testMe_test.go:46: f2(2) != 1
//     testMe_test.go:49: f2(10) != 55
// FAIL
// FAIL    command-line-arguments  0.313s
// FAIL

// ❯ go test testMe.go testMe_test.go -run='F2' -v
// === RUN   TestF2
//     testMe_test.go:43: f2(1) != 1
//     testMe_test.go:46: f2(2) != 1
//     testMe_test.go:49: f2(10) != 55
// --- FAIL: TestF2 (0.00s)
// FAIL
// FAIL    command-line-arguments  0.121s
// FAIL

// ❯ go test testMe.go testMe_test.go -run='F1' -v
// === RUN   TestF1
// --- PASS: TestF1 (0.00s)
// PASS
// ok      command-line-arguments  0.288s

// ❯ go test testMe.go testMe_test.go -run='F1'
// ok      command-line-arguments  0.284s
