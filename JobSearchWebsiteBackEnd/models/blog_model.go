package models

import "time"

type Blog struct {
	ID		uint	`gorm:"primary_key"`
	BlogContext		string
	Email		string
	ReleseTime	time.Time
}
