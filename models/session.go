package models

const SESSION_TABLE = "session"

type Session struct {
	ID         uint
	UUID       string
	Token      string
	UserID     uint
	LastActive string
	ExpiryAt   string
	CreatedAt  string
	UpdatedAt  string
}
