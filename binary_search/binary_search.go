package main

import (
	"fmt"
	"os"
)

func main() {

	//Array must be sorted
	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Println(fmt.Sprintf("Length of array is - %d", len(arr)))
	var a int
	fmt.Print("Enter a number: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println(binary_search(a, arr, len(arr)))

}

func binary_search(a int, array [5]int, len_arr int) int {
	var low, high, middle int

	low = 0
	high = len_arr - 1

	for low <= high {
		middle = (low + high) / 2
		if a < array[middle] {
			high = middle - 1
		} else if a > array[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
