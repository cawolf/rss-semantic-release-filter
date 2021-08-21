package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var logger *zap.SugaredLogger

func main() {
	CreateLogger()

	refresh := flag.NewFlagSet("refresh", flag.ExitOnError)
	refreshDirectory := refresh.String("directory", ".", "directory containing configuration and database")

	feed := flag.NewFlagSet("feed", flag.ExitOnError)
	feedDirectory := feed.String("directory", ".", "directory containing configuration and database")

	if len(os.Args) < 2 {
		PrintUsage()
	}

	switch os.Args[1] {
	case "refresh":
		ParseArgs(refresh)

		RefreshCommand(*refreshDirectory + "/")
	case "feed":
		ParseArgs(feed)

		FeedCommand(*feedDirectory + "/")
	default:
		PrintUsage()
	}
}

func CreateLogger() {
	structuredLoggingEnv := os.Getenv("RSS_SEMANTIC_RELEASE_FILTER_LOG_STRUCTURED")
	structuredLoggingEnabled := false
	if structuredLoggingEnv != "" {
		var err error
		structuredLoggingEnabled, err = strconv.ParseBool(structuredLoggingEnv)
		if err != nil {
			structuredLoggingEnabled = false
		}
	}

	var baseLogger *zap.Logger
	if structuredLoggingEnabled {
		baseLogger, _ = zap.NewProduction()
	} else {
		baseLogger, _ = zap.NewDevelopment()
	}

	defer baseLogger.Sync()
	logger = baseLogger.Sugar()
}

func PrintUsage() {
	fmt.Println("Usage: rss-semantic-release-filter refresh|feed")
	fmt.Println("")
	fmt.Println("    refresh - refreshes all subscribed feeds and updates the internal filtered database")
	fmt.Println("    feed - generates the filtered feed")
	os.Exit(1)
}

func ParseArgs(refresh *flag.FlagSet) {
	err := refresh.Parse(os.Args[2:])
	if err != nil {
		logger.Fatal(err)
	}
}
