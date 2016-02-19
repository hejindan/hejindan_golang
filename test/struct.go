package main

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return *r.width * *r.height
}
func main() {
	var a Rect = Rect{1.4, 4.2, 9.0, 4.0}
	c := a.Area()
	fmt.Println(c)
}
