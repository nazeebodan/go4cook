package main

import (
	"fmt"
	"runtime"
	_ "sync"
)

func mainconcurrent1() {

	c := make(chan bool)
	go func() {
		fmt.Println("go go go!")
		c <- true
		close(c) //如果把这一行注释了，那么就会产生死锁，相当于channel需要一个显式的关闭
	}()
	for v := range c {
		fmt.Println(v)
	}
}
func mainmainconcurrent2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go GoGO(c, i)
	}
	//	<-c
	for i := 0; i < 10; i++ {
		<-c
	}
}

func GoGO(c chan bool, index int) {

	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//	if index == 9 {
	//		c <- true
	//	}
	c <- true
}
