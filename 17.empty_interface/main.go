package main

import "fmt"

type I interface {
}

type S []int

func main() {
	var i I
	fmt.Printf("i Type: %T, Value: %T\n", i, i)

	i = nil
	fmt.Printf("i Type: %T, Value: %T\n", i, i)

	var s S
	fmt.Printf("s Type: %T, Value: %T\n", s, s)

	i = s
	fmt.Printf("i Type: %T, Value: %T\n", i, i)

	var n interface{}
	n = 100
	t, ok := n.(int)
	fmt.Println(t, ok)
}
