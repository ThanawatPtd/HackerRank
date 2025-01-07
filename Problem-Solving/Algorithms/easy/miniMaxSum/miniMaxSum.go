package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func heapify(arr []int, n, i int) {
	largest := i     // Initialize the largest as the root
	left := 2*i + 1  // Left child index
	right := 2*i + 2 // Right child index

	// If the left child is larger than the root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If the right child is larger than the largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If the largest is not the root, swap and continue heapifying
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap
		heapify(arr, n, largest)                    // Recursively heapify the affected subtree
	}
}

// HeapSort sorts an array in ascending order using heap sort.
func heapSort(arr []int) {
	n := len(arr)

	// Build the max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // Move the current root to the end
		heapify(arr, i, 0)              // Heapify the reduced heap
	}
}

func miniMaxSum(arr []int32) {
	cp := make([]int, len(arr))
	for x, _ := range arr {
		cp[x] = int(arr[x])
	}
	heapSort(cp)
	minValue, maxValue := 0, 0
	for i := 0; i < 4; i++ {
		minValue += cp[i]
	}
	for i := 1; i < 5; i++ {
		maxValue += cp[i]
	}
	fmt.Printf("%d %d", minValue, maxValue)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
