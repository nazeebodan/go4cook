package main

import (
	"fmt"
)

func mainDefer() {
	//	testDefer1() //defer类似于堆栈，最后的命令最先执行
	//	testDefer2() //注意闭包，引用的是i的地址！
	orderA()
	orderB()
	orderC()
}

func testDefer1() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func testDefer2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func orderA() {
	fmt.Println("method is orderA")
}

//func orderB() {
//	panic("panic method is orderB")
//}

func orderB() {

	//当遇到异常后，恢复，然后可以继续执行异常之后的程序
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("method is orderB")
		}
	}()
	fmt.Println("???")
	panic("panic method is orderB")
}

func orderC() {
	fmt.Println("method is orderC")
}
