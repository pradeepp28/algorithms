package main

import "fmt"

func merge(first, second []int) []int {
	res := make([]int, len(first)+len(second))
	i, j, k := 0, 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			res[k] = first[i]
			i = i + 1
		} else {
			res[k] = second[j]
			j = j + 1
		}
		k = k + 1
	}
	for ; i < len(first); i++ {
		res[k] = first[i]
		k = k + 1
	}
	for ; j < len(second); j++ {
		res[k] = second[j]
		k = k + 1
	}
	return res
}

func mergeSort(l, h int, nums []int) []int {
	if l == h {
		return nums[l : l+1]
	}
	for l < h {
		mid := l + ((h - l) / 2)
		first := mergeSort(l, mid, nums)
		second := mergeSort(mid+1, h, nums)
		return merge(first, second)
	}
	return []int{}
}

func main() {
	fmt.Println(mergeSort(0, 5, []int{6, 7, 4, 8, 3, 2}))
	fmt.Println(mergeSort(0, 1, []int{3, 2}))
	fmt.Println(mergeSort(0, 1, []int{1, 2}))
}
