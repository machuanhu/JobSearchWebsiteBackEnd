package models

import "time"

type Resume struct {
	ID		uint	`gorm:"primary_key"`
	HrEmail string
	JobSeekerEmail string
	IsReplied	bool
	ReleseTime	time.Time
	ReplyTime	time.Time
	ResumeContext string
	Reply	string
}
