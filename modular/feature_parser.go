package modular

import (
	"strings"
)

func ParseFeatures(input string) []Feature {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	var features []Feature

	var nextLine string
	var skipNext bool
	skipNext = false
	for i, line := range lines {
		if i == len(lines)-1 {
			nextLine = ""
		} else {
			nextLine = lines[i+1]
		}
		if skipNext {
			skipNext = false
			continue
		}
		if i == 0 {
			// First line is always a heading
			features = append(features, Feature{Heading: line, Summary: nextLine})
			skipNext = true
			continue
		}

		if line != "" && nextLine != "" {
			features = append(features, Feature{Heading: line, Summary: nextLine})
			skipNext = true
		} else {
			lastFeature := &features[len(features)-1]
			lastFeature.Summary += "\n" + line
			skipNext = false
		}
	}
	for i := range features {
		features[i].Summary = strings.TrimSpace(features[i].Summary)
	}

	return features
}
