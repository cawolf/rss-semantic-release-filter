package main

import (
	"database/sql"
	"github.com/mmcdole/gofeed"
	_ "modernc.org/sqlite"
)

func OpenDatabase(directory string) *sql.DB {
	db, err := sql.Open("sqlite", "file:"+directory+"feeds.db")
	if err != nil {
		logger.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS items (url TEXT PRIMARY KEY, title TEXT NOT NULL, content TEXT NOT NULL, published INTEGER NOT NULL)")
	if err != nil {
		logger.Fatal(err)
	}

	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_published ON items (published)")
	if err != nil {
		logger.Fatal(err)
	}

	return db
}

func AddItemsToDatabase(db *sql.DB, matchedItems []*gofeed.Item) int64 {
	stmt, _ := db.Prepare("INSERT OR IGNORE INTO items VALUES (?, ?, ?, ?) ")
	rowsAffected := *new(int64)
	for _, item := range matchedItems {
		r, _ := stmt.Exec(item.Link, item.Title, item.Content, item.Published)
		a, _ := r.RowsAffected()
		rowsAffected += a
	}
	logger.Infow("added new items to database",
		"addedItemCount", rowsAffected,
	)
	return rowsAffected
}

func ReadItemsFromDatabase(db *sql.DB) []*gofeed.Item {
	var items []*gofeed.Item
	storedItems, _ := db.Query("SELECT url, title, content, published FROM items ORDER BY published DESC LIMIT 10")

	var (
		url       string
		title     string
		content   string
		published string
	)
	for storedItems.Next() {
		_ = storedItems.Scan(&url, &title, &content, &published)
		items = append(items, &gofeed.Item{
			Title:     title,
			Content:   content,
			Link:      url,
			Published: published,
		})
	}

	_ = storedItems.Close()
	return items
}

func CloseDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.Fatal(err)
	}
}
