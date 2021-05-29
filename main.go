package main

import "fmt"

func main() {
	fmt.Println("yohhh")
	dataTypeAndFVariable()
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
