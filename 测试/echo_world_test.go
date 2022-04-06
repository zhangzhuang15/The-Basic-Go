package main

import (
	"test/utils"
	"testing"
)

func TestEchoWorld(t *testing.T) {
	var result = utils.Echo_World()
	if result != "World" {
		t.Fail()
	}
}
