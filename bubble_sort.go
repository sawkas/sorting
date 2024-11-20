package main

import "fmt"

// Time Complexity: O(N2)
// Auxiliary Space: O(1)

// [0],[9],8,7,6,5,4,3,2,1
// 0,[9],[8],7,6,5,4,3,2,1 => swap
// 0,8,[9],[7],6,5,4,3,2,1 => swap

func BubbleSort(arr []int) {
	var length int = len(arr)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Printf("Array: %d\n", arr)
		}
	}
}

func OptimizedBubbleSort(arr []int) {
	var length int = len(arr)
	var swapped bool

	for i := 0; i < length-1; i++ {
		swapped = false

		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

				swapped = true
			}
			fmt.Printf("Array: %d\n", arr)
		}

		if swapped {
			break
		}
	}
}

func runOptimizedBubbleSort() {
	array := []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("Initial array: %d\n", array)

	OptimizedBubbleSort(array)

	fmt.Printf("Sorted array: %d\n", array)
}

func runBubbleSort() {
	array := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Printf("Initial array: %d\n", array)

	OptimizedBubbleSort(array)

	fmt.Printf("Sorted array: %d\n", array)
}
