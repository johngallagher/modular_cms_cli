package modular

import (
	"reflect"
	"testing"
)

func ParseFeaturesTest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Feature
	}{
		{
			name: "single feature",
			input: `Feature 1
Description 1`,
			expected: []Feature{
				{Heading: "Feature 1", Summary: "Description 1"},
			},
		},
		{
			name: "multiple features",
			input: `Feature 1
Description 1

Feature 2
Description 2`,
			expected: []Feature{
				{Heading: "Feature 1", Summary: "Description 1"},
				{Heading: "Feature 2", Summary: "Description 2"},
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []Feature{},
		},
		{
			name: "extra newlines",
			input: `Feature 1
Description 1


Feature 2
Description 2

`,
			expected: []Feature{
				{Heading: "Feature 1", Summary: "Description 1"},
				{Heading: "Feature 2", Summary: "Description 2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseFeatures(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ParseFeatures() = %v, want %v", got, tt.expected)
			}
		})
	}
}
