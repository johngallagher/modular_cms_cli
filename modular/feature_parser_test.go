package modular

import (
	"reflect"
	"testing"
)

func TestParseFeatures(t *testing.T) {
	t.Run("single feature", func(t *testing.T) {
		input := `Feature 1
Description 1`
		expected := []Feature{
			{Heading: "Feature 1", Summary: "Description 1"},
		}
		got := ParseFeatures(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("ParseFeatures() = %v, want %v", got, expected)
		}
	})

	t.Run("multiple features", func(t *testing.T) {
		input := `Feature 1
Description 1

Feature 2
Description 2`
		expected := []Feature{
			{Heading: "Feature 1", Summary: "Description 1"},
			{Heading: "Feature 2", Summary: "Description 2"},
		}
		got := ParseFeatures(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("ParseFeatures() = %v, want %v", got, expected)
		}
	})

	t.Run("two features with two paragraphs", func(t *testing.T) {
		input := "Feature 1\nDescription 1\n\nParagraph 2\n\nFeature 2\nDescription 2"
		expected := []Feature{
			{Heading: "Feature 1", Summary: "Description 1\n\nParagraph 2"},
			{Heading: "Feature 2", Summary: "Description 2"},
		}
		got := ParseFeatures(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("ParseFeatures() = %v, want %v", got, expected)
		}
	})

	t.Run("extra newlines", func(t *testing.T) {
		input := `Feature 1
Description 1


Feature 2
Description 2

`
		expected := []Feature{
			{Heading: "Feature 1", Summary: "Description 1"},
			{Heading: "Feature 2", Summary: "Description 2"},
		}
		got := ParseFeatures(input)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("ParseFeatures() = %v, want %v", got, expected)
		}
	})
}
