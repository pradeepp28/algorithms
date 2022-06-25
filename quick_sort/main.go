package main

import "fmt"

func quickSort(s []int) {
	if len(s) < 2 {
		return
	}
	i, j := 0, len(s)-1
	pivot := s[i]
	i = i + 1
	for i <= j {

		if s[i] > pivot && s[j] < pivot {
			s[i], s[j] = s[j], s[i]
		}
		if s[i] < pivot {
			i = i + 1
		}
		if s[j] > pivot {
			j = j - 1
		}
	}
	if s[0] > s[j] {
		s[0], s[j] = s[j], s[0]
	}
	quickSort(s[0:j])
	quickSort(s[i:])
}
func main() {
	s := []int{1, 16, 8, 12, 15, 6, 3, 9, 25}
	quickSort(s)
	fmt.Println(s)
	s = []int{1, 2, 3, 5, 6, 7, 8, 9, 25}
	quickSort(s)
	fmt.Println(s)
}
