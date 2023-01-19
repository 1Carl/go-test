package book

import "time"

type Book struct {
	ID            int
	Author        string
	Name          string // Norwegian Wood
	PublishedDate time.Time
	AverageRating float32
	ISBN          string
	Category      string
}
