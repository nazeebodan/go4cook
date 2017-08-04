package main

import (
	"fmt"
)

func mainMethod() {
	//	methodA(1, 2, 3, 4)

	//	methodB()

	//闭包的用法
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

//这个a是一个slice，类似于python的元组
func methodA(a ...int) {
	fmt.Println(a)
}

func methodB() {
	a := func() {
		fmt.Println("这个是匿名函数")
	}
	a()
}

//闭包的用法
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
