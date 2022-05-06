package main

import (
	"math/rand"
	"regexp"
	"time"
)

// Contains 查看切片是否包含某元素
func Contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

// IsBetween 检查给定的时间是否处于某个时间区间
func IsBetween( t , start , end time.Time ) bool {
	if t.After(start) && t.Before(end) {
		return true
	}
	return false
}

// GetSpecificTimestamp 计算特定时区当前时间戳
func GetSpecificTimestamp(timeZone string) time.Time {
	loc, _ := time.LoadLocation(timeZone)

	now := time.Now().In(loc)

	return now
}

// Singlify 删除切片中的重复项
func Singlify(slice []string) []string {
	m := make(map[string]bool)
    s := make([]string, 0)
	for _, value := range slice {
		if _, ok := m[value]; !ok {
			m[value] = true
			s = append(s, value)
		}
	}
	return s
}

// ReverseSlice 切片反转
func ReverseSlice(slice []string) {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
}

// Shuffle 随机打乱切片
func Shuffle(slice []string) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(slice) - 1; i > 0; i-- {
		// 生成一个随机位置，0 ≤ j ≤ i
		j := random.Intn(i+1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Sum 切片求和
func Sum(slice []int) int {
	sum := 0

	for _, value := range slice {
		sum += value
	}

	return sum
} 

// ConvertToSnakeCase 字符串转换为 snake 书写形式
func ConvertToSnakeCase(s string) string {
	matchChars := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAlpha := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchChars.ReplaceAllString(s, "_")
	snake = matchAlpha.ReplaceAllString(snake, "_")

	return snake
}


func main() {

}