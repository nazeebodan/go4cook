package main

import (
	"fmt"
)

type human struct {
	sex int
}
type person struct {
	name string
	age  int
	human
	contact struct {
		phone, address string
	}
}

func main() {
	structInit2()
	fmt.Println("==================")
	structInit3()
}

func structInit1() {
	s := person{}
	s.name = "suredandan"
	s.age = 18
	fmt.Println(s)
}

//虽然没有class的概念，没有构造函数的概念，但是有defer，所以就类似了
func structInit2() {
	s := person{
		name: "nazibodan",
		age:  18,
	}
	fmt.Println(s)
	modify(s) //结构体是值引用，虽然在函数里面改变了，但是是改变的拷贝
	fmt.Println(s)
}

func structInit3() {
	//结构体是值引用，所以需要用&来对一个改变生效
	s := &person{
		name:  "nazibodan",
		age:   18,
		human: human{sex: 100}, //这个是嵌入字段的初始化
	}
	//	*s.name //go语言中如果是地址引用的，不需要这么做
	s.name = "nazi wave"
	fmt.Println(s)
	modify2(s)
	s.contact.address = "你猜"
	s.contact.phone = "9090950"
	fmt.Println(s)
	var b = &person{}
	b = s
	b.contact.phone = "8888888"
	b.sex = 1
	fmt.Println(b)
}

func modify(per person) person {
	per.age = 30
	fmt.Println(per)
	return per
}

func modify2(per *person) *person {
	per.age = 30
	fmt.Println(per)
	return per
}
