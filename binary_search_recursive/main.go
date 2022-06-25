package main

import "fmt"

func search(arr []int, i, j, key int) int {
	if i == j {
		if arr[i] == key {
			return i
		} else {
			return -1
		}
	}
	mid := (i + j) / 2
	if key == arr[mid] {
		return mid
	} else if key < arr[mid] {
		return search(arr, 0, mid, key)
	} else {
		return search(arr, mid+1, j, key)
	}
}
func main() {
	arr := []int{3, 6, 8, 12}
	fmt.Println(search(arr, 0, len(arr)-1, 12)) // 3
	fmt.Println(search(arr, 0, len(arr)-1, 3))  // 0
	fmt.Println(search(arr, 0, len(arr)-1, 6))  // 1
	fmt.Println(search(arr, 0, len(arr)-1, 8))  // 2
	fmt.Println(search(arr, 0, len(arr)-1, 30)) // -1
	fmt.Println(search(arr, 0, len(arr)-1, 1))  // -1

}
