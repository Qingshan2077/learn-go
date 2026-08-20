package ch3

import "fmt"

var u uint8 = 255

func ShowUint8() {
	fmt.Println(u, u+1, u*u)
	// "255 0 1"
}

var i int8 = 127

func Showint8() {
	fmt.Println(i, i+1, i*i) // "127 -128 1"
}
