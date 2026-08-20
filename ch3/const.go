package ch3

// 常见的常量
const IPv4Len = 4

// iota用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。
// 在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
