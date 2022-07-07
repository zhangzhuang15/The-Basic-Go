package main

import (
	"fmt"
)

// 回溯算法架构
// track(当前结果集, 选择列表) {
//    if 当前结果集符合要求 {
//      res.push(当前结果集)
//      return
//    }
//
//    for 选择 in 选择列表 {
//      if 选择 not in 当前结果集 {
//	         当前结果集.push(选择)
//           track(当前结果集, 新的选择列表)
//           当前结果集.pop(选择)
//      }
//    }
// }
// 这里特别感谢 labuladong 的算法小抄，深入浅出，非常帮助小白.


func hasValue(temp []int, value int) bool {
	for _, item := range temp {
		if item == value {
			return true
		}
	}
	return false
}
// 以全排列为例
// 输入 [1, 2, 3]
// 输出 [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1]]
func getResult(input []int) [][]int {
	res := make([][]int, 0)

	var backtrack func(temp []int)
	backtrack = func(temp []int) {
		if len(temp) == len(input) {
			_temp := append([]int{}, temp...)
			res = append(res, _temp)
			return
		}
        // NOTE: for index, value := range vector
		// 不要理解为 for value := range vector
		for _, item := range input {
			if hasValue(temp, item) {
				continue
			}

			temp = append(temp, item)
			backtrack(temp)
			temp = temp[0: len(temp) - 1]
		}
	}
	backtrack([]int{})
    return res
}

func main() {
    input := []int {3, 4, 5}
	result := getResult(input)
	fmt.Print("result: ", result)
}