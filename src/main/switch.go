package main

import (
	"fmt"
)

func switch1() {
	fmt.Println("switch1")
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a!=0")

	}
}

func switch2() {
	fmt.Println("switch2")
	a := 1
	switch {
	case a > 0:
		fmt.Println("a>0")
		fallthrough
	case a == 1:
		fmt.Println("a=1")

	}
}

func switch3() {
	fmt.Println("switch3")
	switch a := 1; {
	case a > 0:
		fmt.Println("a>0")
		fallthrough
	case a == 1:
		fmt.Println("a=1")

	}
}

//func main() {
//	switch3()
//}
