package ch2

import "fmt"

// ShowFor 一个简单的for循环实现逆序遍历
func ShowFor() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}
