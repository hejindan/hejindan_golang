package main

import "fmt"

type Vertex struct {
	X int
	Y int
	P string
}

func main() {
	var a Vertex = Vertex{1, 2, "hjd"}
	var c Vertex = a
	fmt.Println(c)
	c.X = 100
	fmt.Println(a)

}
