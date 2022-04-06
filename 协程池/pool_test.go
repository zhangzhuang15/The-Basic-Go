package pool

import (
	"fmt"
	"testing"
)

func feibo(value int) int {
	if value == 0 || value == 1 {
		return value
	}
	return feibo(value-1) + feibo(value-2)
}

func TestRun10Tasks(t *testing.T) {
	pool := New(4, 10)
	for i := 0; i < 10; i++ {
		pool.AddTask("Jack")
	}
	if pool.taskSize != 10 {
		t.Fail()
	}
	var worker func(task interface{})
	worker = func(task interface{}) {
		t := task.(string)
		feibo(42)
		fmt.Println(t)
	}
	pool.Run(worker)
	pool.Exit()
	if pool.taskSize != 0 {
		t.Fail()
	}
}

func TestRun100Tasks(t *testing.T) {
	pool := New(8, 100)
	for i := 0 ; i < 100; i++ {
		pool.AddTask("Tom")
	}
	var worker func(task interface{})
	worker = func(task interface{}) {
		t := task.(string)
		feibo(40)
		fmt.Println(t)
	}
	pool.Run(worker)
	pool.Exit()
	if pool.taskSize != 0 {
		t.Fail()
	}
}
