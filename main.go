package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	refresh := flag.NewFlagSet("refresh", flag.ExitOnError)
	feed := flag.NewFlagSet("feed", flag.ExitOnError)

	if len(os.Args) < 2 {
		PrintUsage()
	}

	switch os.Args[1] {
	case "refresh":
		ParseArgs(refresh)

		RefreshCommand()
	case "feed":
		ParseArgs(feed)

		FeedCommand()
	default:
		PrintUsage()
	}
}

func ParseArgs(refresh *flag.FlagSet) {
	err := refresh.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
}

func PrintUsage() {
	fmt.Println("Usage: rss-semantic-release-filter refresh|feed")
	fmt.Println("")
	fmt.Println("    refresh - refreshes all subscribed feeds and updates the internal filtered database")
	fmt.Println("    feed - generates the filtered feed")
	os.Exit(1)
}
