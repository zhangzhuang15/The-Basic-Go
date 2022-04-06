package maxpool

// MaxPool 对array进行窗口大小为group，步长为slide的最大池化操作
func MaxPool(array []int, group, slide int) []int {
	if group <= 0 || slide <= 0 || len(array) == 0 {
		return []int{}
	}

	// 存储结果
	result := make([]int, 0)

	// 这个情况下，没有办法借助上一个窗口的结果优化运算
	if group <= slide {
		for left, right := 0, group-1; right < len(array); left, right = left+slide, right+slide {
			maxValue := array[left]
			// 当前窗口搜索最大值
			for i := left; i <= right; i++ {
				if array[i] > maxValue {
					maxValue = array[i]
				}
			}
			result = append(result, maxValue)
		}

		return result
	}

	// 双边队列, 统计最大值的情况
	deq := make([]int, 0)

	for i := 0; i < group-1; i++ {
		if len(deq) == 0 {
			deq = append(deq, i)
		} else {
			// 在deq中，从结尾到开头的方向搜索第一个满足 array[deq[j]] > array[i],
			// 之后将 deq[j+1] 更新为 i， deq[j+1]后边的值全部删除（也就是出队列）
			j := len(deq) - 1
			for ; j >= 0; j-- {
				if array[i] < array[deq[j]] {
					break
				}
			}
			if j == -1 {
				deq[0] = i
				deq = deq[0:1]
			}
			if j == len(deq) - 1 {
				deq = append(deq, i)
			}else {
				deq[j+1] = i
				deq = deq[0: j+2]
			}
		}
	}

	// i 指向每一个窗口的结尾元素
	for i := group-1; i < len(array); i += slide {
		for k := i-slide+1 ; k <= i ; k++ {
			if k < group-1 {
				continue
			}
			j := 0
			for ; j < len(deq); j++ {
				// deq[j] 指向的元素，不在窗口内，过期失效
				if i - deq[j] >= group {
					continue
				}
				// 在没过期的数据中，找到了以往的最大值 array[deq[j]],
				// 当前元素 array[k] 比这个值大的话，更新 deq
				if array[deq[j]] < array[k] {
					deq[0] = k
					deq = deq[0:1]
				} else {
					// 以往最大值 array[deq[j]]仍旧是最大值, 之后要将 array[k] 通过比较纳入到 deq 中。
					p := len(deq) - 1
					for ; p >= j ; p-- {
						if array[deq[p]] > array[k] {
							break
						}
					}
					if p == -1 {
						deq[0] = k
						deq = deq[0:1]
					}
					if p == len(deq) - 1 {
						deq = append(deq, k)
						deq = deq[j:]
					} else {
						deq[p+1] = k
						deq = deq[j: p+2]
					}
				}
				break
			}
		}
		result = append(result, array[deq[0]])
	}

	return result
}
