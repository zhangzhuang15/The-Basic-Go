package seqsum


// 问题描述：
// 给你一个数组arr[10], 每个元素的初始值为0，
// 给你一个二维数组 conditions = { {1, 4, 10}, { 3, 9, -9}, {2, 6, 11}},
// {1, 4, 10} 的意思是令 arr[1] arr[2] arr[3] arr[4] 的元素都加上 10。
//
// 按照condition提供的规则，计算arr最终的结果

// 通常的解决方式，O(n^2)
func sumSlow(arr []int, conditions [][]int) []int {
	for _, condition := range conditions {
		start, end, value := condition[0], condition[1], condition[2]
		for i := start; i <= end; i++ {
			arr[i] += value
		}
	}
	return arr
}

// 累加和方法，O(n)
// 例子： arr[5] = [0, 0, 0, 0, 0]
//       conditions = [ [1, 3, 2]]
//       只需要令 arr[1] += 2, arr[4] -= 2
//       于是 arr[5] = [0, 2, 0, 0, -2]
//       之后采取累加和 arr[i+1] = arr[i+1] + arr[i]
//       得到 arr[5] = [0, 2, 2, 2, 0]
//       很巧妙的想法，对吧？
func sumFast(arr []int, conditions [][]int) []int {
	for _, condition := range conditions {
		start, end, value := condition[0], condition[1], condition[2]
		arr[start] += value
		if end < len(arr) - 1{
			arr[end + 1] -= value
		}
	}

	for i := 0; i < len(arr) - 1; i++ {
		arr[i+1] += arr[i]
	}

	return arr
}
