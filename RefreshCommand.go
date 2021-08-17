package main

import (
	"github.com/mmcdole/gofeed"
	"log"
	"regexp"
)

func RefreshCommand() {
	const SemanticVersioningRegexpPattern = "(?P<major>0|[1-9]\\d*)\\.(?P<minor>0|[1-9]\\d*)\\.(?P<patch>0|[1-9]\\d*)(?:-(?P<prerelease>(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?"

	semanticVersioningRegexp, _ := regexp.Compile(SemanticVersioningRegexpPattern)

	fp := gofeed.NewParser()

	db := OpenDatabase()

	configuration := ReadConfiguration()

	log.Println("refreshing all subscribed feeds") // TODO: show number of feeds

	// TODO: loop
	feed := GetFeed(*fp) // TODO use configuration: feedUrl
	// TODO: handle nil feed

	filteredFeedItems := FilterFeed(feed, semanticVersioningRegexp, configuration)

	addedItemCount := AddItemsToDatabase(db, filteredFeedItems)

	CloseDatabase(db)

	log.Printf("refreshed all subscribed feeds, found %d items matching the filters, added %d items to database\n", len(filteredFeedItems), addedItemCount) // TODO: show number of feeds
}
