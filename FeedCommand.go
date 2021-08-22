package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type dbContext struct {
	echo.Context
	db *sql.DB
}

func FeedCommand(directory string) {
	db := OpenDatabase(directory)
	defer CloseDatabase(db)

	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &dbContext{c, db}
			return h(cc)
		}
	})

	e.GET("/", ServeFeed)

	e.Logger.Fatal(e.Start(":1323"))
}

func ServeFeed(c echo.Context) error {
	cc := c.(*dbContext)

	items := ReadItemsFromDatabase(cc.db)

	feed := GenerateFeed(items)

	feedString, _ := feed.ToRss()

	return c.Blob(http.StatusOK, echo.MIMEApplicationXMLCharsetUTF8, []byte(feedString))
}
