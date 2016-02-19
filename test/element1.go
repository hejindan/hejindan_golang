package main

import "fmt"

func modify(array []int) {
	array[0] = 10
	fmt.Println("In modify(),array values:", array)
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("In main(),array values:", array)
	array1 := array[:5]
	modify(array1)
	fmt.Println("In main(),array1 values:", array1)
	fmt.Println("In main(),array1 values:", array)

}
