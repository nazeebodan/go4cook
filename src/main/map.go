package main

import (
	"fmt"
	"sort"
)

func main2() {
	//	var m map[int]string = make(map[int]string)
	testmap4()
}

func testmap1() {
	m := make(map[int]string)
	m[1] = "ok"
	m[3] = "hha"
	fmt.Println(m[1])
	delete(m, 1)
	fmt.Println(m[1])
}

func testmap2() {
	m := make(map[int]map[int]string)

	//复杂map的赋值，要先去检查里层的是否赋值了
	a, ok := m[2][1]
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "nice"
	a, ok = m[2][1]
	fmt.Println(a)
}

func testmap3() {
	//以map为元素的slice
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "oooo"
	}
	fmt.Println(sm)

	for i, _ := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "oooo"
	}
	fmt.Println(sm)
}

func testmap4() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m)
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

}
