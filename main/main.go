package main

import (
	search "algorithms/algorithms/search"
	"fmt"
)

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(search.BinarySearch(arr, 6))
}
