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
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

// GetString reads floats from a file and stores them in a slice
func GetString(fileName string) ([]string, error) {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
