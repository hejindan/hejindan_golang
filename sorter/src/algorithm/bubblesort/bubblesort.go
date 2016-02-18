package bubblesort

func BubbleSort(values []int) {
	for i := 0; i < len(values); i++ {
		for j := i; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
