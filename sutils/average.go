// Package sutils uses slices for all its computations
package sutils

import "strconv"

// Average function uses slices to calcuate the average
func Average(nums []int) int {
	var total int
	for _, value := range nums {
		total += value
	}

	return total / len(nums)
}

// InputAvg get average of all the numbers entered from the command line
func InputAvg(nums []string) (float64, error) {
	var total float64
	for _, num := range nums {
		number, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return total, err
		}
		total += number
	}
	return total / float64(len(nums)), nil
}
