package typeTest

import (
	"fmt"
)

// PrintAll
/**
any:等价于 interface{},任何类型都可以用 any 表示,这个类型的变量可以存放任意类型的值
可变参数: ...any 表示可以接受任意数量的任意类型参数
*/
func PrintAll(args ...any) {
	fmt.Println(args)
}
