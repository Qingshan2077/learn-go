package ch3

import "fmt"

func ShowString() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])                 // "104 119" ('h' and 'w')
	fmt.Println(string(s[0]), string(s[7])) // 'h' and 'w'
	fmt.Println(s[0:5])                     // "hello"
	// go的字符串是不可修改的
	a := "left foot"
	t := a
	a += ", right foot"
	fmt.Println(s) // "left foot, right foot"
	fmt.Println(t) // "left foot"

}
