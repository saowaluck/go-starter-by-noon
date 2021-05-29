package main

import (
	"fmt"
)

func main() {
	fmt.Println("yohhh")
	dataTypeAndFVariable()
	loop()
	condition()
	arrayandSlide()
}

func dataTypeAndFVariable() {
	// static type
	var s string
	var i int
	var f float64
	var b bool

	// zero
	fmt.Println(s, i, f, b)

	// assign
	s = "hi"
	i = 1
	f = 1.0
	b = false
	fmt.Println(s, i, f, b)

	// short declear variable
	s1 := "hi"
	i1 := 1
	f1 := 1.0
	b1 := false
	fmt.Println(s1, i1, f1, b1)

	// constance
	const pi float64 = 3.17
	fmt.Println(pi)
}

func loop() {
	// increment loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for with condition
	i := 1
	for i < 5 {
		i++
		fmt.Println(i)
	}

	// for ever (true)
	for {
		fmt.Println("loop")
		if true {
			break
		}
	}

	// for each
	numbers := []int{1, 2}
	for i, v := range numbers {
		fmt.Println(i, v)
	}
}

func condition() {
	// if else
	if 1 == 2 {
		fmt.Println("1 == 2")
	} else {
		fmt.Println("1 !== 2")
	}

	// switch case
	switch 1 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("noting")
	}
}

func arrayandSlide() {
	// array (storage memory, fix length)
	var x = [2]int{1, 2}
	x[1] = 3
	fmt.Println(x)

	// slice
	var y = []int{}
	y = append(y, 0)
	fmt.Println(y)
}
