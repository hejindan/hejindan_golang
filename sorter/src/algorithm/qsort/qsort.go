package qsort

import "fmt"

func quicksort(values []int, left, right int) {
	temp := left
	if left == right {
		return
	}
	for i := left; i <= right; i++ {
		if values[temp] > values[i] {
			//	values[temp], values[i] = values[i], values[temp]
			//	temp = i
			k := values[i]
			for j := i; j > temp; j-- {
				values[j] = values[j-1]
			}
			values[temp] = k
			temp = temp + 1
		}
		fmt.Printf("i:%d\n", i)
		fmt.Printf("temp:%d\n", temp)
	}
	if temp != 0 {
		quicksort(values, 0, temp-1)
	}
	if temp != right {
		quicksort(values, temp+1, right)
	}
}

func QuickSort(values []int) {
	quicksort(values, 0, len(values)-1)
}
