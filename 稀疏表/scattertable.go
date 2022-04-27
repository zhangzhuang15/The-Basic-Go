package scattertable


// 问题：
//    给你一个数组 arr[5] = [ 1, 3, 4, 2, 7]
//    让你求解 arr[2:4]的最大值，最小值，和？
//
//    区间问题，可以用线段树解决，但这种问题可以不用线段树，
//    why？
//    因为不需要对数组 arr 进行动态操作啊。
//
//    那可以用什么方法？
//    稀疏表

// MaxIn 返回 arr[start: end] 的最大值（不包括arr[end]）
//
// start: 起始索引  0 <= start < len(arr)
//
// end :  结束索引  0 < end <= len(arr)
func MaxIn(arr []int, start, end int) int {
	// scatterTable[i][j] 表示 arr[i: i+2^j] 的最大值,
	// scatterTable[i][j] = max{ scatterTable[i, j-1], scatterTable[ i+2^(j-1), j-1] }
	scatterTable := make([][]int, len(arr))

	// scatterTable[i][j]
	// 当 i = 0 时，对应的右边界为索引 0 + 2^j - 1,
	// 应该有 2^j - 1 < len(arr),
	// 即 2^j <= len(arr)
	maxJ := 0
	for ; (1 << maxJ) <= len(arr); maxJ++ {}

	for i := 0; i < len(arr); i++ {
		scatterTable[i] = make([]int, maxJ)
		scatterTable[i][0] = arr[i]
	}

	for i := 0; i < len(arr); i++ {
		// scatterTable[i][j] 对应的是 arr[i: i+ 2^j]
		// 因此 i + 2^j <= len(arr), 即 i + 2^j -1 < len(arr)
		for j := 1; j < maxJ && i + (1<<j) - 1 < len(arr); j++ {
			scatterTable[i][j] = max(scatterTable[i][j-1], scatterTable[i + (1 << (j-1))][j-1])
		}
	}

	j := 0
	// 寻找满足 start + 2^j - 1 <= end -1 条件的最大j值
	// 该条件等效于 2^j <= end - start,
	// 进而等效于 2^j < end - start + 1，
	// 当循环结束时，条件没有的到满足，那么 j-1 就是满足条件的最大值了
	for ; (1 << j) < (end - start + 1); j++ {}
    j -= 1

	return max(scatterTable[start][j-1], scatterTable[end-(1<<(j-1))][j-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 指举出了求区间最大值的场景，
// 对于别的场景，
// 只需要更新 稀疏表scatterTable的递推关系式，并修改相关代码即可

// 由于 arr 是不会变化的，
// 在第一次求解的时候会花费 O(n^2)时间，
// 但在之后的求解中，scatterTable不需要重新建立，时间只需花费在45～53行，只需要 O(logn)时间。

// 当然，本例子中只给出了第一次求解的场景，
// 并没有做缓存处理，
// 这一点上，代码仍可优化，
// 但不妨碍我们领会稀疏表的动态规划算法思想。