package qsort

import "fmt"

func quicksort(values []int, left, right int) {
	temp := left
	//if left == right {
	//	return
	//}
	for i := left; i <= right; i++ {
		if values[temp] > values[i] {
			values[temp], values[i] = values[i], values[temp]
			temp = i
		}
		fmt.Printf("i:%d\n", i)
		fmt.Printf("temp:%d\n", temp)
	}
	//quicksort(values, 0, temp-1)
	//quicksort(values, temp+1, right)
}

func QuickSort(values []int) {
	quicksort(values, 0, len(values)-1)
}
