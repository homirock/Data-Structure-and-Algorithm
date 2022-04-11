package main

import (
	"fmt"
)
//insertion sort.
func insertionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
func main() {
	arr := []int{7, 8, 4, 10, 1}
	insertionSort(arr)
	fmt.Println(insertionSort(arr))
}
