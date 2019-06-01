package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part %s) - indexes: expecting %d <> found %d"

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Go is cool", "Go", 0},
		{"", "", 0},
		{"Hey", "hey", -1},
		{"Matheus", "t", 2},
		{"Matheus", "t", 0},
	}

	for _, test := range tests {
		t.Logf("Mass: %v", test)
		cur := strings.Index(test.text, test.part)

		if cur != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, cur)
		}
	}
}
