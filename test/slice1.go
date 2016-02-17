package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{7, 8, 9}
	for i := 0; i < len(slice1); i++ {
		fmt.Println("slice1[", i, "] =", slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Println("slice2[", i, "] =", slice2[i])
	}
	copy(slice2, slice1)
	fmt.Println("after copy\n")
	for i := 0; i < len(slice1); i++ {
		fmt.Println("slice1[", i, "] =", slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Println("slice2[", i, "] =", slice2[i])
	}

}
