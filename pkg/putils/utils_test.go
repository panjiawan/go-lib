package putils

import (
	"fmt"
	"testing"
)

func TestUtils(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 2, 3, 4}

	// 映射：将数字转换为字符串
	strings := Map(numbers, func(n int) string {
		return fmt.Sprintf("数字%d", n)
	})
	fmt.Println("映射结果:", strings)

	// 过滤：只保留偶数
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("偶数:", evens)

	// 归约：求和
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("总和:", sum)

	// 包含判断
	fmt.Println("包含3:", Contains(numbers, 3))
	fmt.Println("包含10:", Contains(numbers, 10))

	// 去重
	uniqueNumbers := Unique(numbers)
	fmt.Println("去重后:", uniqueNumbers)
}
