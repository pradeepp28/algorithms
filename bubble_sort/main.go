package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- { //num of comparision
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func main() {
	arr := []int{4, 2, 1, 5, 9, 7}

	bubbleSort(arr)
	fmt.Println(arr)
	arr = []int{1, 5, 4, 7, 2, 9}

	bubbleSort(arr)
	fmt.Println(arr)
}
