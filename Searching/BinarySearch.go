package main

import "fmt"

func BinarySearch(arr []int, num int) bool {
	min, max := 0, len(arr)-1
	for min<max {
		mid := (min + (max-min)) / 2
		if arr[mid] == num {
			return true
		} else if arr[mid] < num {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return false
}
func main() {
	array := []int{1,2,3,4}
	fmt.Println(BinarySearch(array,5))
}
