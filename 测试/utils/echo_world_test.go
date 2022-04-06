package utils

import "testing"

func TestEchoWorld2(t *testing.T) {
	var result = Echo_World()
	if result != "world" {
		t.Fail()
	}
}
