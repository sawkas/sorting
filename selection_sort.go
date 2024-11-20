package main

import "fmt"

// Time Complexity: O(N2)
// Auxiliary Space: O(1)

// [0],9,8,7,6,1,2,3,4,[5] => no swap
// 0,[9],8,7,6,[1],2,3,4,5 => swap
// 0,1,[8],7,6,9,[2],3,4,5 => swap

func SelectionSort(arr []int) {
	length := len(arr)
	var minIndex int

	for i := 0; i < length; i++ {
		minIndex = i

		for j := i; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

		fmt.Printf("Array: %d\n", arr)
	}
}

func StableSelectionSort(arr []int) {
	// TODO:
}

func runSelectionSort() {
	array := []int{0, 9, 8, 7, 6, 1, 2, 3, 4, 5}

	fmt.Printf("Initial array: %d\n", array)

	SelectionSort(array)

	fmt.Printf("Sorted array: %d\n", array)
}
