package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(2, 4)
	if r != 6 {
		t.Fatalf("add(2,4) error, execpt:%d, actual:%d", 6, r)

	}
	t.Logf("test add success")

}

func TestSub(t *testing.T) {
	r := Sub(4, 2)
	if r != 2 {
		t.Fatalf("add(2,4) error, execpt:%d, actual:%d", 6, r)

	}
	t.Logf("test add success")

}
