package main

import (
	"fmt"
	"github.com/Danil-114195722/mygomodules/arrays"
)

func main() {
	fmt.Println("Hello, world!")

	arr1 := []int{1, 2, 3}
	value1 := 3
	value2 := 6

	fmt.Printf("value1 in arr1: %v\n", arrays.IntInArray(arr1, value1))
	fmt.Printf("value2 in arr1: %v\n", arrays.IntInArray(arr1, value2))
}
