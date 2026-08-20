// 打印水的沸点
package ch2

//导包
import "fmt"

// 常量定义
const boilingF = 212.0

// Boiling 首字母大写，才能在别的包中被引用，首字母大写就相当于别的语言中的public字段
func Boiling() {
	//var f float64 = boilingF
	//var f = boilingF
	f := boilingF
	//var c = (f - 32) * 5 / 9
	fmt.Printf("沸点为 = %g°F or %g°C\n", f, fTOC(f))
}

// 调用含参函数
func fTOC(f float64) float64 {
	return (f - 32) * 5 / 9
}
