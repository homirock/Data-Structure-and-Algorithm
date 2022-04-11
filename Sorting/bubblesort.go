package main

import (
	"fmt"
)
// bubblesort.
func bubblesort(arr []int) []int {
	len := len(arr)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{10, 5, 4, 8, 9}
	bubblesort(arr)
	fmt.Println(arr)
}
