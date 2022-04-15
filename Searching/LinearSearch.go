package main

import "fmt"

func LinearSearch(arr []int, n int) (bool, error) {

	l := len(arr)
	for i := 0; i < l; i++ {
		if n == arr[i] {
			return true, nil
		}
	}
	return false, fmt.Errorf("Element is not present in the array")

}
func main() {
	arr := []int {5,7,2,3,1}
	fmt.Println(LinearSearch(arr,56))
}