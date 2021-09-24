package models

import "time"

type Comment struct {
	ID		uint	`gorm:"primary_key"`
	Email	string
	BlogId	int
	CommentContext	string
	ReleseTime	time.Time
}