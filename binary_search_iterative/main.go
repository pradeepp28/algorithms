package main

import "fmt"

func search(arr []int, key int) int {
	i, j := 0, len(arr)
	for i < j {
		mid := (i + j) / 2
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return -1
}
func main() {
	// arr := []int{3, 6, 8, 12, 14, 17, 25, 29, 31, 36, 42, 47, 53, 55, 62}
	arr := []int{3, 6, 8, 12}
	fmt.Println(search(arr, 12)) // 3
	fmt.Println(search(arr, 3))  // 0
	fmt.Println(search(arr, 6))  // 1
	fmt.Println(search(arr, 8))  // 2
	fmt.Println(search(arr, 30)) // -1
	fmt.Println(search(arr, 1))  // -1
}
