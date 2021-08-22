package main

import (
	"github.com/Masterminds/semver/v3"
	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
	"regexp"
	"time"
)

func GetFeed(parser gofeed.Parser, feedConfiguration FeedConfiguration) *gofeed.Feed {
	feed, _ := parser.ParseURL(feedConfiguration.Url)

	if feed != nil {
		logger.Infow("refreshing feed",
			"feedTitle", feed.Title,
		)
	} else {
		logger.Warnw("could not fetch feed, skipping",
			"feedUrl", feedConfiguration.Url,
		)
	}
	return feed
}

func FilterFeed(feed *gofeed.Feed, regexp *regexp.Regexp, feedConfiguration FeedConfiguration) []*gofeed.Item {
	var filteredFeedItems []*gofeed.Item

	for _, item := range feed.Items {
		if IsVersionMatchingTheFilter(item.Title, regexp, feedConfiguration.MinimumLevel) {
			filteredFeedItems = append(
				filteredFeedItems,
				item,
			)
		}
	}

	logger.Infow("filtered matching feed items",
		"filteredFeedItemCount", len(filteredFeedItems),
	)
	return filteredFeedItems
}

func IsVersionMatchingTheFilter(versionString string, regexp *regexp.Regexp, level MinimumLevelType) bool {
	version, _ := semver.NewVersion(regexp.FindString(versionString))

	if level == Major && (version.Minor() != 0 || version.Patch() != 0 || version.Prerelease() != "") {
		return false
	}

	if level == Minor && (version.Patch() != 0 || version.Prerelease() != "") {
		return false
	}

	if level == Patch && version.Prerelease() != "" {
		return false
	}

	return true
}

func GenerateFeed(items []*gofeed.Item) *feeds.Feed {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       "filtered semantic releases",
		Link:        &feeds.Link{Href: "https://github.com/cawolf/rss-semantic-release-filter"},
		Description: "RSS feeds of projects filtered by semantic version levels",
		Author:      &feeds.Author{Name: "Christian A. Wolf", Email: "mail@cawolf.de"},
		Created:     now,
	}

	var feedItems []*feeds.Item

	for _, item := range items {
		itemTime, _ := time.Parse(time.RFC3339, item.Published)
		feedItem := &feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Content,
			Created:     itemTime,
		}
		feedItems = append(feedItems, feedItem)
	}

	feed.Items = feedItems

	return feed
}
