package main

import (
	"github.com/mmcdole/gofeed"
	"regexp"
)

const SemanticVersioningRegexpPattern = "(?P<major>0|[1-9]\\d*)\\.(?P<minor>0|[1-9]\\d*)\\.(?P<patch>0|[1-9]\\d*)(?:-(?P<prerelease>(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?"

func RefreshCommand(directory string) {
	semanticVersioningRegexp, _ := regexp.Compile(SemanticVersioningRegexpPattern)

	fp := gofeed.NewParser()

	db := OpenDatabase(directory)

	var configuration Configuration
	configuration.Read(directory)

	logger.Infow("refreshing subscribed feeds",
		"feedCount", len(configuration.Feeds),
	)

	successfullyRefreshed := 0
	filteredFeedItemCount := 0
	addedFeedItemCount := *new(int64)
	for _, feedConfiguration := range configuration.Feeds {
		feed := GetFeed(*fp, feedConfiguration)

		if feed == nil {
			continue
		}

		filteredFeedItems := FilterFeed(feed, semanticVersioningRegexp, feedConfiguration)

		filteredFeedItemCount = filteredFeedItemCount + len(filteredFeedItems)

		addedFeedItemCount = addedFeedItemCount + AddItemsToDatabase(db, filteredFeedItems)

		successfullyRefreshed++
	}

	CloseDatabase(db)

	logger.Infow("refreshed subscribed feeds",
		"successfullyRefreshedFeeds", successfullyRefreshed,
		"filteredFeedItemCount", filteredFeedItemCount,
		"addedFeedItemCount", addedFeedItemCount,
	)
}
