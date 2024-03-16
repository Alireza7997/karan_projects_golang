package classic_algorithms

import (
	"fmt"
	"strings"
)

// Sorting :
// Implement two types of sorting algorithms: Merge sort and bubble sort.
func Sorting() {
	var (
		availableSorts []string = []string{"\033[32mBubble(b)", "Merge(m)\033[0m"}
		sortType       string
		array          []int
		sortedArray    []int
	)

	fmt.Printf("Enter sorting type: %v\n", availableSorts)
	fmt.Scanln(&sortType)
	sortType = strings.ToLower(sortType)

	switch sortType {
	case "bubble", "b":
		array = getArray()
		sortedArray = bubbleSort(array)
		fmt.Printf("Original array: %v\n", array)
		fmt.Printf("Sorted array: %v\n", sortedArray)

	case "merge", "m":
		array = getArray()
		sortedArray = mergeSort(array)
		fmt.Printf("Original array: %v\n", array)
		fmt.Printf("Sorted array: %v\n", sortedArray)

	default:
		fmt.Printf("This sorting type is not available : \033[31m%s\033[0m \n", sortType)
		Sorting()
	}

}
func getArray() []int {
	var (
		length int
	)
	fmt.Print("Enter the length of the array: ")
	fmt.Scanln(&length)
	array := make([]int, length)

	for i := 0; i < length; i++ {
		var number int
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanln(&number)
		array[i] = number
	}

	return array
}
func bubbleSort(arr []int) (sorted []int) {
	length := len(arr)
	if length <= 1 {
		sorted = arr
		return
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	sorted = arr

	return
}

func mergeSort(arr []int) (sorted []int) {
	if len(arr) <= 1 {
		sorted = arr
		return
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	sorted = merge(left, right)
	return
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
