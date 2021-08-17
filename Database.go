package main

import (
	"database/sql"
	"github.com/mmcdole/gofeed"
	"log"
	_ "modernc.org/sqlite"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "file:./foo.db") // FIXME
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS items (url TEXT PRIMARY KEY, title TEXT NOT NULL, content TEXT NOT NULL, published INTEGER NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_published ON items (published)")
	if err != nil {
		log.Fatal(err)
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
	log.Printf("added %d items to database", rowsAffected)
	return rowsAffected
}

func CloseDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
