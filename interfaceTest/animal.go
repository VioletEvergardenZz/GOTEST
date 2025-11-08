// Package Animal
/**
接口只是一种“契约”。它是一种仅声明方法而不实现它们的类型。
结构体定义了属性和方法。
可以将结构体类型的变量传入期望接口类型参数的函数中。
只要结构体类型实现了接口提供的"契约"，它就可以作为接口的实例使用。
类型检查系统会验证结构体是否定义了接口声明的方法，如果满足条件，类型检查就会允许它像接口实例一样被使用或传递
*/
package Animal

import (
	"fmt"
)

// Animal 定义两个结构体将实现的接口
type Animal interface {
	Speak() string
	Move() string // string表示该方法返回一个字符串类型的值，描述动物的移动方式
}

// Dog  Dog struct 是一种具有一个字段的类型，但它们通常具有许多字段
type Dog struct {
	Name string
}

// Cat 为Cat定义一个类似的结构体。注意，这里的字段不需要和狗一样。接口对这些字段一无所知。
type Cat struct {
	Name string
}

// Speak 实现 for Dog，匹配接口中的相同函数
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!\n", d.Name)
}

// Move Dog 的移动实现，匹配接口中的相同函数
func (d Dog) Move() string {
	return fmt.Sprintf("%s runs around happily.\n", d.Name)
}
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!\n", c.Name)
}
func (c Cat) Move() string {
	return fmt.Sprintf("%s gracefully walks away.\n", c.Name)
}

// DescribeAnimal 接收接口作为参数的函数，换句话说，任何定义了Speak和Move方法的struct
func DescribeAnimal(a Animal) {
	fmt.Println(a.Speak())
	fmt.Println(a.Move())
	fmt.Println()
}
