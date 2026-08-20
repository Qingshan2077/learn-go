package ch2

import "fmt"

func Gcd(x, y int) {
	fmt.Printf("%d 和 %d 的最大公约数是：%d\n", x, y, calculate(x, y))
}

func calculate(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
