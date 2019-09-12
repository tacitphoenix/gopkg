// Package keyboard reads user input from the keyboard
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Input reads an int from the keyboard
// It returns the number read and any error encountered
func Input() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	} else {
		input = strings.TrimSpace(input)
	}

	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return value, nil
}
