package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MergeSort initializes the auxiliary slice and starts the sorting process.
// This ensures we only allocate memory once, adhering strictly to O(N) space complexity.
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// Pre-allocate a single auxiliary slice of the same size
	aux := make([]int, len(arr))

	// Copy the initial unordered data into the auxiliary slice
	copy(aux, arr)

	// Start the recursive sort
	sort(arr, aux, 0, len(arr)-1)
}

// sort recursively divides the slice and merges the sorted halves.
// We swap the roles of 'arr' and 'aux' at each level to avoid unnecessary copying.
func sort(arr, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2

	// Sort the left and right halves recursively
	// Notice that the arguments for arr and aux are swapped
	sort(aux, arr, lo, mid)
	sort(aux, arr, mid+1, hi)

	// Merge the sorted halves back together
	merge(arr, aux, lo, mid, hi)
}

// merge takes two sorted sub-arrays (inside 'aux') and merges them into 'arr'.
func merge(arr, aux []int, lo, mid, hi int) {
	i := lo      // Starting index for the left sub-array
	j := mid + 1 // Starting index for the right sub-array

	for k := lo; k <= hi; k++ {
		if i > mid {
			// Left sub-array is exhausted
			arr[k] = aux[j]
			j++
		} else if j > hi {
			// Right sub-array is exhausted
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			// Right element is smaller than left element
			arr[k] = aux[j]
			j++
		} else {
			// Left element is smaller than or equal to right element
			arr[k] = aux[i]
			i++
		}
	}
}

func main() {
	const numElements = 10000

	// 1. Generate 10,000 pseudo-random integers
	data := make([]int, numElements)

	// Seed the random number generator to ensure different sequences on each run
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := 0; i < numElements; i++ {
		data[i] = r.Intn(100000) // Generate numbers between 0 and 99,999
	}

	fmt.Printf("Generated a slice of %d unordered elements.\n", numElements)
	fmt.Printf("Sample (first 5 unsorted): %v\n", data[:5])

	// 2. Perform the sort
	start := time.Now()
	MergeSort(data)
	duration := time.Since(start)

	// 3. Validate the sort
	isSorted := true
	for i := 1; i < numElements; i++ {
		if data[i-1] > data[i] {
			isSorted = false
			break
		}
	}

	// 4. Output the results
	if isSorted {
		fmt.Printf("\nSuccess! The slice was sorted correctly in %v.\n", duration)
		fmt.Printf("Sample (first 5 sorted): %v\n", data[:5])
		fmt.Printf("Sample (last 5 sorted): %v\n", data[numElements-5:])
	} else {
		fmt.Println("\nError: The sort failed. The data is out of order.")
	}
}
