package main

import (
	"GOTEST/interfaceTest"
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
	dog := Animal.Dog{Name: "Buddy"}
	cat := Animal.Cat{Name: "Misty"}
	// 两个结构都满足 Animal 接口
	Animal.DescribeAnimal(dog)
	Animal.DescribeAnimal(cat)
}
