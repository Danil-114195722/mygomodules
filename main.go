package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("This source contain modules for Go.")
	fmt.Println("List of modules use prefix <github.com/Danil-114195722/mygomodules/> to import module:")
	fmt.Printf("\n\t arrays \t This module contain funcs for arrays\n")
	fmt.Printf("\n\t\t IntInArray(arr []int, value int) -- Check \"value\" (type int) in array \"arr\" (type []int)\n")
	fmt.Printf("\n\t\t StringInArray(arr []string, value string) -- Check \"value\" (type string) in array \"arr\" (type []string)\n")
    fmt.Printf("\n\t\t FloatInArray(arr []float64, value float64) -- Check \"value\" (type float64) in array \"arr\" (type []float64)\n")
}
