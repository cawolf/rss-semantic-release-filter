package main

import (
	"github.com/Masterminds/semver/v3"
	"github.com/mmcdole/gofeed"
	"regexp"
)

func GetFeed(parser gofeed.Parser, feedConfiguration FeedConfiguration) *gofeed.Feed {
	feed, _ := parser.ParseURL(feedConfiguration.FeedUrl)

	if feed != nil {
		logger.Infow("refreshing feed",
			"feedTitle", feed.Title,
		)
	} else {
		logger.Warnw("could not fetch feed, skipping",
			"feedUrl", feedConfiguration.FeedUrl,
		)
	}
	return feed
}

func FilterFeed(feed *gofeed.Feed, regexp *regexp.Regexp, feedConfiguration FeedConfiguration) []*gofeed.Item {
	var filteredFeedItems []*gofeed.Item

	for _, item := range feed.Items {
		version, _ := semver.NewVersion(regexp.FindString(item.Title))

		if feedConfiguration.ComparisonLevel == Major && (version.Minor() != 0 || version.Patch() != 0 || version.Prerelease() != "") {
			continue
		}

		if feedConfiguration.ComparisonLevel == Minor && (version.Patch() != 0 || version.Prerelease() != "") {
			continue
		}

		if feedConfiguration.ComparisonLevel == Patch && version.Prerelease() != "" {
			continue
		}

		filteredFeedItems = append(
			filteredFeedItems,
			item,
		)
	}

	logger.Infow("filtered matching feed items",
		"filteredFeedItemCount", len(filteredFeedItems),
	)
	return filteredFeedItems
}
