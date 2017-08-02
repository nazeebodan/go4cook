package main

import (
	"fmt"
)

//func main() {
//
//	//	for1()
//	//	for2()
//	for3()
//	a := 1
//	var p *int = &a
//	fmt.Println(p)
//	fmt.Println(*p)
//}

/*
	第一种for循环的写法，for后面不带任何参数，相当于while(true)
**/
func for1() {
	a := 1
	for {
		fmt.Println(a)
		a++
		if a > 3 {
			break
		}
	}
}

/*
	第二张for循环的写法
**/
func for2() {
	a := 1
	for a <= 3 {
		fmt.Println(a)
		a++
	}
}

/*
	for的第三种写法，在for后面并行的加一些初始化语句
*/
func for3() {
	a := 1
	for i := 0; i < 3; i++ {
		fmt.Println(a)
		a++
	}

}
