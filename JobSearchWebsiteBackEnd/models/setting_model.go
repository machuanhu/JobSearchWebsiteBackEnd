package models


type Config struct{
	Key			string			`json:"key"`
	DbSettings	DbSettings		`json:"DbSettings"`
	HostName	string			`json:"HostName"`
}

type  DbSettings struct {
	Username	string	`json:"Username"`
	Password	string	`json:"Password"`
	Hostname	string	`json:"Hostname"`
	Dbname		string	`json:"Dbname"`
}