package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	array1 := array[3:]
	array2 := array[:3]
	fmt.Println("in main,array,values:", array)
	fmt.Println("in main,array,values:", array1)
	fmt.Println("in main,array,values:", array2)
}
