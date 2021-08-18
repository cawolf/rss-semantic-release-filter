package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func FeedCommand(directory string) {
	// TODO

	logger.Info("feed")

	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	/*fp := gofeed.NewParser()
	feed := GetFeed(*fp)

	return c.XML(http.StatusOK, feed.Items[0].Published+": "+feed.Items[0].Title+" "+feed.Items[0].Link+"\n"+feed.Items[0].Content)*/
	return nil
}
