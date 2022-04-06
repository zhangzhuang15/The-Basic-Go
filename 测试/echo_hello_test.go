package main

import "testing"

func TestEchoHello(t *testing.T) {

	var result = EchoHello()
	if result != "hello" {
		t.FailNow()
	}
}
