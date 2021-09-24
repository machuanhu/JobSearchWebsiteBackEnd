package models

type User struct {
	Email             string `gorm:"primary_key"`
	Username          string
	Password          string
	RegisterTimestamp int64
	Role              string
	Introduction	  string
}
