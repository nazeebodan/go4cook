package main

import (
	"fmt"
	"runtime"
	"sync"
)

func mainc2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	//	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go GoTest(&wg, i)
	}
	wg.Wait()

}

func GoTest(wg *sync.WaitGroup, index int) {

	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
