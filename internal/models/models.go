package models

import "time"

type Item struct {
	Title      string
	Categories []string
	Link       string
	Date       time.Time
	Summary    string
	SourceName string
}

type Source struct {
	ID          int64
	Name        string
	URL         string
	CreatedDate time.Time
}

type Article struct {
	ID            int64
	SourceID      int64
	Title         string
	Link          string
	Summary       string
	PublishedDate time.Time
	PostedDate    time.Time
	CreatedDate   time.Time
}
