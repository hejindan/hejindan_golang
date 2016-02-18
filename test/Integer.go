package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) Integer {
	a = a + b
	return a + b
}
func main() {
	var a Integer = 1
	c := a.Less(2)
	fmt.Println(c)

}
