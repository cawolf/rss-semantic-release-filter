package main

import (
	"github.com/mmcdole/gofeed"
	"regexp"
)

func RefreshCommand(directory string) {
	const SemanticVersioningRegexpPattern = "(?P<major>0|[1-9]\\d*)\\.(?P<minor>0|[1-9]\\d*)\\.(?P<patch>0|[1-9]\\d*)(?:-(?P<prerelease>(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?"

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
			logger.Warnw("could not fetch feed, skipping",
				"feedUrl", feedConfiguration.FeedUrl,
			)
			continue
		}

		filteredFeedItems := FilterFeed(feed, semanticVersioningRegexp, feedConfiguration)

		filteredFeedItemCount = filteredFeedItemCount + len(filteredFeedItems)

		addedFeedItemCount = addedFeedItemCount + AddItemsToDatabase(db, filteredFeedItems)

		CloseDatabase(db)

		successfullyRefreshed++
	}

	logger.Infow("refreshed subscribed feeds",
		"successfullyRefreshedFeeds", successfullyRefreshed,
		"filteredFeedItemCount", filteredFeedItemCount,
		"addedFeedItemCount", addedFeedItemCount,
	)
}
