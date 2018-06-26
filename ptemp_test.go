package main

import (
	"testing"

	"github.com/coderconvoy/lz2"
)

func Test_PTemp(t *testing.T) {
	pt, err := NewPTemp(`hello {{join "poo" .}}`, &lz2.Config{})

	if err != nil {
		t.Fail()
	}

	r, err := pt.Exec("world")

	if err != nil {
		t.Fail()
	}

	if string(r) != "hello poo/world" {
		t.Errorf("got : %s", r)
	}
}
