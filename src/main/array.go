package main

import (
	"fmt"
)

//func main() {
//	a := [2]int{2}
//	fmt.Println(a)
//	b := [20]int{19: 1}
//	fmt.Println(b)
//	c := [...]int{1, 2, 3, 4, 5}
//	fmt.Println(len(c))
//	d := [...]int{1: 2}
//	fmt.Println(d)
//}

/**
指针数组
*/
//func main() {
//	x, y := 1, 2
//	a := [...]*int{&x, &y}
//	fmt.Println(a)
//}

/*
数组指针
*/
func main1() {
	a := [...]int{99: 1}
	var p *[100]int = &a
	fmt.Println(p)

	p2 := new([100]int)
	fmt.Println(p2)  //指向数组的指针
	fmt.Println(*p2) //数组
	fmt.Println(&p2) //数组的地址

	//============================
	//直接给数组赋值或者通过指针来赋值，效果是一样的
	b := [10]int{}
	b[1] = 2
	fmt.Println(b)

	p3 := new([10]int)
	p3[1] = 2
	fmt.Println(p3)
	//======================
	sortTest()
}

func sortTest() {
	a := [...]int{3, 2, 6, 5, 1, 4}
	fmt.Println(a)

	cnt := len(a)
	m, n := 0, 0

	for i := 0; i < cnt; i++ {
		m++
		for j := i + 1; j < cnt; j++ {
			n++
			if a[i] > a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
	fmt.Println(a)
	fmt.Println(m)
	fmt.Println(n)

}
