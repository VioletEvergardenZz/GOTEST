package typeTest

import (
	"fmt"
)

//Wrap
/**
wrap 是一个高阶函数，它的参数和返回值都是函数
参数: fn - 一个接受任意数量任意类型参数并返回任意类型值的函数
返回值: 一个接受任意数量任意类型参数但无返回值的函数
*/
func Wrap(fn func(...any) any) func(...any) {
	// 定义一个内部函数
	fnInner := func(params ...any) {
		res := fn(params...) //把切片拆成多个独立参数
		fmt.Println(res)     //打印出 fn 的返回结果
	}
	return fnInner
}
