package main

import (
	"alghoritms/alghoritms/sortings"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sourceSlice := createSlice()

	copiedSlice1 := copySlice(sourceSlice)
	copiedSlice2 := copySlice(sourceSlice)
	copiedSlice3 := copySlice(sourceSlice)
	copiedSlice4 := copySlice(sourceSlice)

	fmt.Printf("BubbleSort: %v\n", showTime(copiedSlice1, sortings.BubbleSort))
	fmt.Printf("SelectionSort: %v\n", showTime(copiedSlice2, sortings.SelectionSort))
	fmt.Printf("InsertionSort: %v\n", showTime(copiedSlice3, sortings.InsertionSort))
	fmt.Printf("QuickSort: %v\n", showTime(copiedSlice4, sortings.QuickSort))
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func createSlice() []int {
	slice := make([]int, 10000)
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func showTime(slice []int, sortingFunc func(slice []int) []int) time.Duration {
	start := time.Now()
	sortingFunc(slice)
	return time.Since(start)
}
