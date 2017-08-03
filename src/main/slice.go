package main

import (
	"fmt"
	_ "os"
	_ "runtime"
)

func main3() {
	//	fmt.Println(runtime.NumCPU())
	test4()

}

func test1() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:9] //包含起始索引不包含终止索引
	fmt.Println(s1)
	s2 := a[5:]
	fmt.Println(s2)
	s3 := a[5:len(a)]
	fmt.Println(s3)

	//一般初始化用make来进行=============
	s4 := make([]int, 10, 10)
	fmt.Println(s4)
	fmt.Println(len(s4))
}

func test2() {
	a := [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println(cap(a), a)

	sa := a[3:7]
	fmt.Println(sa)
	sb := sa[3:5] //slice 的本质是指向数组的连续数组，所以它的最大cap是起始索引到数组的长度，和cap有关，所以用len和cap可以看出来结果为什么会这样
	fmt.Println(sb)
	fmt.Println(len(sa), cap(sa)) //sa的cap=a的cap-sa的起始索引=10-3=7
	fmt.Println(len(sb), cap(sb)) //sb的cap=sa的cap-sb的起始索引=7-3=4
}

func test3() {
	s1 := make([]int, 3, 6)
	s2 := append(s1, 1, 2, 3)
	s3 := append(s1, 1, 2, 3, 4)
	fmt.Println(cap(s1))
	fmt.Println(cap(s2)) //如果没有超过s1的cap，那么仍然是6，所以cap(s2)=6
	fmt.Println(cap(s3)) //如果超过了s1的cap，那么扩容就按照s1的cap来进行扩容，所以cap(s3)=12，换句话说内存地址也会发生变化，用fmt.Println("%p\n",sX)可以查看
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", s3)

}

func test4() {
	a := []int{1, 2, 3, 4, 5, 6}
	s1 := a[1:3]
	s2 := a[2:4]
	fmt.Println(s1, s2)
	s1[1] = 0 //如果多个slice同时指向同一个数组，那么任意改变共有的值，就全部都会改变
	fmt.Println(s1, s2)
}
