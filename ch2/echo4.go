package ch2

import (
	"flag"
	"fmt"
	"strings"
)

// 变量声明：
// 简短变量声明被广泛用于大部分的局部变量的声明和初始化 如i := 100  等价于 var i int =100
// var形式的声明语句往往是用于需要显式指定变量类型的地方 如var names []string

// 指针：
//x := 1
//p := &x         // p, of type *int, points to x
//fmt.Println(*p) // "1"
//*p = 2          // equivalent to x = 2
//fmt.Println(x)  // "2"

// Echo4 prints its command-line arguments.

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func Echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
