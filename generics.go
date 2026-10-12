package main

import "fmt"

func generics[T any](value []T) {
	for _, va := range value {
		fmt.Print(va)
	}
}

func main() {
	val1 := []int{1, 2, 3}
	generics(val1)
}
