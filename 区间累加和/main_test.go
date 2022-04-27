package seqsum

import (
	"testing"
	"time"
)

func TestSumSlow(t *testing.T) {
	arr := make([]int, 2000)
	conditions := [][]int { {1, 4, 2}, { 3, 7, -2}, {0, 2, 1} }

	// 计算复杂化 conditions的规模，sumSlow仍旧执行非常快，
	// 伤脑筋了，没办法体现它和 sumFast的性能差距。
    start := time.Now()
	result := sumSlow(arr, conditions)
    end := time.Now().UnixMilli() - start.UnixMilli()
    t.Logf("cost time: %d ms\n", end)

	_result := []int {1, 3, 3, 0, 0, -2, -2, -2}

	for i := 0; i < len(_result); i++ {
		if result[i] != _result[i] {
			t.Fail()
		}
	}
}

func TestSumFast(t *testing.T) {
	arr := []int {0 , 0, 0, 0, 0, 0, 0, 0}
	conditions := [][]int { {1, 4, 2}, { 3, 7, -2}, {0, 2, 1} }

	start := time.Now()
	result := sumFast(arr, conditions)
	end := time.Now().UnixMilli() - start.UnixMilli()
	t.Logf("cost time: %d ms\n", end)

	_result := []int {1, 3, 3, 0, 0, -2, -2, -2}

	for i := 0; i < len(_result); i++ {
		if result[i] != _result[i] {
			t.Fail()
		}
	}
}
