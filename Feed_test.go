package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestFilterFeed(t *testing.T) {
	compiledRegexp, _ := regexp.Compile(SemanticVersioningRegexpPattern)

	tests := []struct {
		version  string
		level    MinimumLevelType
		expected bool
	}{
		{version: "v1.0.0", level: Patch, expected: true},
		{version: "v1.0.0", level: Minor, expected: true},
		{version: "v1.0.0", level: Major, expected: true},
		{version: "v1.1.0", level: Patch, expected: true},
		{version: "v1.1.0", level: Minor, expected: true},
		{version: "v1.1.0", level: Major, expected: false},
		{version: "v1.1.1", level: Patch, expected: true},
		{version: "v1.1.1", level: Minor, expected: false},
		{version: "v1.1.1", level: Major, expected: false},
		{version: "v1.0.0+b1043", level: Patch, expected: true},
		{version: "v1.0.0+b1043", level: Minor, expected: true},
		{version: "v1.0.0+b1043", level: Major, expected: true},
		{version: "v1.1.0+b1043", level: Patch, expected: true},
		{version: "v1.1.0+b1043", level: Minor, expected: true},
		{version: "v1.1.0+b1043", level: Major, expected: false},
		{version: "v1.1.1+b1043", level: Patch, expected: true},
		{version: "v1.1.1+b1043", level: Minor, expected: false},
		{version: "v1.1.1+b1043", level: Major, expected: false},
		{version: "v1.0.0-rc", level: Patch, expected: false},
		{version: "v1.0.0-rc", level: Minor, expected: false},
		{version: "v1.0.0-rc", level: Major, expected: false},
		{version: "v1.1.0-rc", level: Patch, expected: false},
		{version: "v1.1.0-rc", level: Minor, expected: false},
		{version: "v1.1.0-rc", level: Major, expected: false},
		{version: "v1.1.1-rc", level: Patch, expected: false},
		{version: "v1.1.1-rc", level: Minor, expected: false},
		{version: "v1.1.1-rc", level: Major, expected: false},
		{version: "v1.0.0-rc+b1043", level: Patch, expected: false},
		{version: "v1.0.0-rc+b1043", level: Minor, expected: false},
		{version: "v1.0.0-rc+b1043", level: Major, expected: false},
		{version: "v1.1.0-rc+b1043", level: Patch, expected: false},
		{version: "v1.1.0-rc+b1043", level: Minor, expected: false},
		{version: "v1.1.0-rc+b1043", level: Major, expected: false},
		{version: "v1.1.1-rc+b1043", level: Patch, expected: false},
		{version: "v1.1.1-rc+b1043", level: Minor, expected: false},
		{version: "v1.1.1-rc+b1043", level: Major, expected: false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%s/%s", test.version, test.level)
		t.Run(testName, func(t *testing.T) {
			assert.EqualValues(t, test.expected, IsVersionMatchingTheFilter(test.version, compiledRegexp, test.level))
		})
	}
}
