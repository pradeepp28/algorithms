package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] > arr[j-1] {
				break
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
func main() {
	arr := []int{4, 2, 1, 5, 9, 7}

	insertionSort(arr)
	fmt.Println(arr)
}
