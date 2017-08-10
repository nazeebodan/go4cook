package main

import (
	"fmt"
)

type USB interface {
	name() string
	Conn
}
type Conn interface {
	connect()
}

type phoneConnecter struct {
	insertName string
}

func (pc phoneConnecter) name() string {
	return pc.insertName
}

func (pc phoneConnecter) connect() {
	fmt.Println("connect:", pc.insertName)
}

func disConnect2(usb USB) {
	if pc, ok := usb.(phoneConnecter); ok {
		fmt.Println(pc.name() + " disconnect!")
		return
	}
	fmt.Println("unknow device!")
}

func disConnect(usb interface{}) {
	//	if pc, ok := usb.(phoneConnecter); ok {
	//		fmt.Println(pc.name() + " disconnect!")
	//		return
	//	}
	switch v := usb.(type) {
	case phoneConnecter:
		fmt.Println(v.name() + " disconnect!")
	default:
		fmt.Println("unknow device!")
	}
}

func main() {
	//	var a USB
	//	a = phoneConnecter{insertName: "nazibodan"}
	a := phoneConnecter{insertName: "nazibodan"}
	a.connect()
	disConnect(a) //因为phoneConnecter实现了USB接口的2个方法，所以相当于就是继承了这个接口了，所以这个地方可以直接传入a
}
