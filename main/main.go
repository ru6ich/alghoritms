package main

import (
	search "algorithms/algorithms/search"
	"fmt"
)

func main() {
	arr := []int{1, 3, 4, 5, 7, 8, 9}
	fmt.Println(search.BinarySearch(arr, 1))
}
