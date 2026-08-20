package ch2

import "fmt"

// range 迭代时，返回每个位置的索引和值。
func ShowRange() {
	// 定义一个 map
	userAges := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Eve":   35,
	}

	// 遍历 map，获取 key 和 value
	for name, age := range userAges {
		fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
	}

	// 只获取 key
	for name := range userAges {
		fmt.Printf("姓名: %s\n", name)
	}

	// 只获取 value
	for _, age := range userAges {
		fmt.Printf("年龄: %d\n", age)
	}
}
