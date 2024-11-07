package models

import "time"

type Project struct {
	Name        string
	Description string
	UsedSkills  string
}

type Skill struct {
	Name        string
	Description string
	Level       string
}

type Experience struct {
	Organisation     string
	Responsibilities string
	DateStart        time.Time
	DateEnd          time.Time
}

type Owner struct {
	FirstName string
	LastName  string
	BirthDate string
	About     string
}

type Contacts struct {
	PhoneNumber string
	Email       string
	Telegram    string
	Linkedin    string
	GitHub      string
}
