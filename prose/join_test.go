package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list     []string
	expected string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, expected: ""},
		testData{list: []string{"apple"}, expected: "apple"},
		testData{list: []string{"apple", "orange"}, expected: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, expected: "apple, orange, and pear"},
	}

	for _, test := range tests {
		actual := JoinWithCommas(test.list)
		if actual != test.expected {
			t.Errorf(errString(test.list, actual, test.expected))
		}
	}
}

func errString(list []string, actual string, expected string) string {
	return fmt.Sprintf("Actual \"%s\" Expected \"%s\"", actual, expected)
}
