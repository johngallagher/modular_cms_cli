package modular

import (
	"bufio"
	"strings"
)

// ParseFeatures converts a string containing feature definitions into []Feature
// Format expected:
// Feature Name 1
// Description for feature 1
//
// Feature Name 2
// Description for feature 2
func ParseFeatures(input string) []Feature {
	features := []Feature{}
	var currentFeature Feature

	scanner := bufio.NewScanner(strings.NewReader(input))
	isName := true // alternates between name and description

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Empty line marks the end of a feature
		if line == "" {
			if currentFeature.Heading != "" {
				features = append(features, currentFeature)
				currentFeature = Feature{}
			}
			isName = true
			continue
		}

		if isName {
			currentFeature.Heading = line
			isName = false
		} else {
			currentFeature.Summary = line
			isName = true
		}
	}

	// Add the last feature if exists
	if currentFeature.Heading != "" {
		features = append(features, currentFeature)
	}

	return features
}
