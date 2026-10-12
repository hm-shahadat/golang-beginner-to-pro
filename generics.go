package main

import "fmt"

func generics[T any](value []T) {
	for _, va := range value {
		fmt.Println(va)
	}
}

//using generics in struct
type stack[T any] struct {
	elements []T
}

func main() {
	val1 := []int{1, 2, 3}
	val2 := []string{"shahadat", "gazi"}
	val3 := []bool{true, false}
	generics(val2)
	generics(val1)
	generics(val3)

	//stack
	myStack := stack[int]{
		elements: []int{4, 5, 6},
	}
	myStack2 := stack[string]{
		elements: []string{"Hossain"},
	}
	fmt.Println(myStack)
	fmt.Println(myStack2)
}
