package main

type Configuration struct {
	feedUrl         string
	comparisonLevel ComparisonLevelType
}

type ComparisonLevelType string

const (
	Major ComparisonLevelType = "major"
	Minor ComparisonLevelType = "minor"
	Patch ComparisonLevelType = "patch"
)

func ReadConfiguration(directory string) Configuration {
	// TODO read configuration from file
	configuration := Configuration{
		feedUrl:         "dsajfdsalsf",
		comparisonLevel: Patch,
	}

	// TODO validate configuration
	return configuration
}
