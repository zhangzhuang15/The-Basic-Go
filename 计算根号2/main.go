package main

import (
	"fmt"
	"math"
)

func isEqual(a, b float64) bool {
	if math.Abs(b - a) < 1e-7 {
		return true
	}
	return false
}

// 二分法
func sqrtMethodOne(value float64) float64 {
	start, end := 0.0, value
	mid := (start + end)/2

	for !isEqual(mid * mid, value) {
		if mid * mid < value {
			start = mid
		} else {
			end = mid
		}
		mid = (start + end)/2
	}

	return mid
}

// 牛顿迭代法
// 曲线：f(x) = x^2 - n
// 曲线的切率： f'(x) = 2x
// 曲线(x0, f(x0))的切线： y(x) = f'(x0)(x-x0) + f(x0)
// 令 y(x) = 0 得到 x1 = x0 - f(x0)/f'(x0)
// 完成了由 x0 迭代到 x1,
// 接下来继续如此迭代，当迭代到 f(xn) = 0 时，切线为 y(x) = f'(xn)(x-xn),
// 此时令 y(x) = 0 得到 xm = xn - f(xn)/f'(xn) = xn - 0/f'(xn) = xn,
// 这样就找到了 f(x) = 0 的零点， 也就是 x = sqrt(n),
// 当然不可能百分百确保 xm == xn, 只要 xm - xn 的绝对值非常小，就可以认为
// xm == xn 了
func sqrtMethodTwo(value float64) float64 {
	// 曲线函数
	f := func(x float64) float64 {
		return x*x - value
	}
	// 曲线的切率
	df := func(x float64) float64 {
		return 2*x
	}
	// x0：上次迭代的点
	// x1：本次迭代的点
	x0, x1 := 0.0, value

	for !isEqual(x0, x1) {
		// 本次迭代结束，x1的值变成x0的值
		x0 = x1
		// 重新计算，开启下一次迭代
		x1 = x0 - f(x0)/df(x0)
	}
	return x1
}

func main() {
	fmt.Println("sqrt_1 (2): ", sqrtMethodOne(2))
	fmt.Println("sqrt_1 (5): ", sqrtMethodOne(5))
	fmt.Println("sqrt_1 (17): ", sqrtMethodOne(17))

	fmt.Println("sqrt_2 (2): ", sqrtMethodTwo(2))
	fmt.Println("sqrt_2 (5): ", sqrtMethodTwo(5))
	fmt.Println("sqrt_2 (17): ", sqrtMethodTwo(17))
}