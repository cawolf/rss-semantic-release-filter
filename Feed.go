package main

import (
	"github.com/Masterminds/semver/v3"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"regexp"
)

func GetFeed(parser gofeed.Parser) *gofeed.Feed {
	file, _ := os.Open("/home/cwolf/Downloads/releases.atom") // TODO: refactor to fetching from github, ...
	defer file.Close()

	feed, _ := parser.Parse(file)

	log.Printf("refreshing '%s'", feed.Title)
	return feed
}

func FilterFeed(feed *gofeed.Feed, regexp *regexp.Regexp, configuration Configuration) []*gofeed.Item {
	var filteredFeedItems []*gofeed.Item

	for _, item := range feed.Items {
		version, _ := semver.NewVersion(regexp.FindString(item.Title))

		if configuration.comparisonLevel == Major && (version.Minor() != 0 || version.Patch() != 0 || version.Prerelease() != "") {
			continue
		}

		if configuration.comparisonLevel == Minor && (version.Patch() != 0 || version.Prerelease() != "") {
			continue
		}

		if configuration.comparisonLevel == Patch && version.Prerelease() != "" {
			continue
		}

		filteredFeedItems = append(
			filteredFeedItems,
			item,
		)
	}

	log.Printf("found %d items matching the filters", len(filteredFeedItems))
	return filteredFeedItems
}
