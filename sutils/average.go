// Package sutils uses slices for all its computations
package sutils

// Average function uses slices to calcuate the average
func Average(nums []int) int {
	var total int
	for _, value := range nums {
		total += value
	}

	return total / len(nums)
}
