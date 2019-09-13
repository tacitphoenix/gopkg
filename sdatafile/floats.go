// Package sdatafile uses slices to read floats from a file
package sdatafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads floats from a file and stores them in a slice
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64

	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}
