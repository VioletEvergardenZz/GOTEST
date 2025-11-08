package main

import (
	"GOTEST/typeTest"
)

func main() {
	/**
	"GOTEST/chanTest"
	*/
	//dfs.Dfs()
	//dfs.DfsDemo()

	/**
	"GOTEST/chanTest"
	*/
	//chanTest.ChanTest1()
	//chanTest.ChanTest2()

	/**
	"GOTEST/interfaceTest"
	*/
	//dog := Animal.Dog{Name: "Buddy"}
	//cat := Animal.Cat{Name: "Misty"}
	//// 两个结构都满足 Animal 接口
	//Animal.DescribeAnimal(dog)
	//Animal.DescribeAnimal(cat)

	/**
	"GOTEST/typeTest"
	*/
	//any0.go
	//typeTest.PrintAll("hello", 1, true)

	/**
	"GOTEST/typeTest"
	*/
	//any.go
	add := func(args ...any) any {
		a := args[0].(int)
		b := args[1].(int)
		return a + b
	}
	wrapedAdd := typeTest.Wrap(add) //fnInner(),返回值类型 func(...any),记录了它要调用的 fn（也就是 add）
	wrapedAdd(3, 4)                 //fnInner 真正被执行,params := []any{3, 4}
}
