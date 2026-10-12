package main

import "fmt"

func generics[T any, v string](value []T, name v) {
	for _, va := range value {
		fmt.Println(va, name)
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
	generics(val2, "aiub")
	generics(val1, "iub")
	generics(val3, "nsu")

	//for struct
	myStack := stack[int]{
		elements: []int{4, 5, 6},
	}
	myStack2 := stack[string]{
		elements: []string{"Hossain"},
	}
	myStack3 := stack[bool]{
		elements: []bool{true},
	}
	fmt.Println(myStack)
	fmt.Println(myStack2)
	fmt.Println(myStack3)
}
