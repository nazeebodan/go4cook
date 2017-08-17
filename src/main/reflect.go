package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id       int
	UserName string
	Age      int
}

func (u User) Hello() {
	fmt.Println("hello,world!")
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("just 4 interface struct")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface() //go语言里面struct里面变量如果大写则是public,如果是小写则是private的，private的时候通过反射不能获取其值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	fmt.Println("methods:")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type) //如果method的首字母为小写，那么反射也获取不到
	}
}

func mainreflect() {
	u := User{1, "nazibo", 18}
	info(&u) //info(u) 可以查看区别，所以才有那个异常的捕获
}
