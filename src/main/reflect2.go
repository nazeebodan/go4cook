package main

import (
	"fmt"
	"reflect"
)

type UserModel struct {
	Id   int
	Name string
	Age  int
}

func (u UserModel) DoWhat(what string) {
	fmt.Println(u.Name + " 在 " + what)
}
func mainreflect2() {
	u := UserModel{1, "suredandan", 18}
	fmt.Println(u)
	Set(&u)
	fmt.Println(u)
	u.DoWhat("swim")
	testReflectMethod(&u)
}

func testReflectMethod(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() != reflect.Ptr { //reflect.Ptr表示是指针类型
		fmt.Println("不能修改！")
	}
	mv := v.MethodByName("DoWhat")
	args := []reflect.Value{reflect.ValueOf("ppp")}
	mv.Call(args)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() != reflect.Ptr && !v.Elem().CanSet() { //reflect.Ptr表示是指针类型
		fmt.Println("不能修改！")
	} else {
		v = v.Elem()
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("nazibo")
	}

	f := v.FieldByName("othetFieldName")
	if !f.IsValid() {
		fmt.Println("aaaggg")
	}
	f = v.FieldByName("Age")
	if !f.IsValid() {
		fmt.Println("bbbggg")
	}
}
