package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "重複があるスライス",
			input:    []string{"dummy1", "dummy2", "dummy2"},
			expected: []string{"dummy1", "dummy2"},
		},
		{
			name:     "重複がないスライス",
			input:    []string{"dummy1", "dummy2"},
			expected: []string{"dummy1", "dummy2"},
		},
		{
			name:     "全ての要素が同じスライス",
			input:    []string{"dummy", "dummy", "dummy", "dummy"},
			expected: []string{"dummy"},
		},
		{
			name:     "空のスライス",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RemoveDuplicates(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
