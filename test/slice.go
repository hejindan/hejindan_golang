package main
import "fmt"
func main() {
	myarray := [10]int{1,2,3,4,5,6,7,8,9,10}
	var myslice []int = myarray[5:]
	fmt.Println("elements of myarray: ")
	for _,v := range myarray{
		fmt.Print(v," ")
	}
	fmt.Println("\nelements of myslice: ")
	for _,v := range myslice{
		fmt.Print(v," ")
	}
	fmt.Println()

}
