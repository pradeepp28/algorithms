package main

import "fmt"

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

//nolint
func heapSortRecursive(arr []int, n int) {
	if n < 2 {
		return
	}
	arr[0], arr[n-1] = arr[n-1], arr[0]
	heapify(arr, n-1, 0)
	// fmt.Println(arr)
	heapSortRecursive(arr, n-1)
}

func main() {
	arr := []int{4, 2, 1, 5, 7, 9}

	heapSort(arr)
	fmt.Println(arr)

	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	heapSortRecursive(arr, len(arr))
	fmt.Println(arr)
}
