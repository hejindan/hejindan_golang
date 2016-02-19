package main

import "fmt"

//func change(a *int)int{
//	*a=10
//	return 0
//}
func change(a *int) int {
	*a = 10
	return 0
}
func main() {
	v3 := 5
	fmt.Printf("hello world")
	fmt.Println(v3)
	change(&v3)
	fmt.Println(v3)

}
