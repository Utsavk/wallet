package models

type Session struct {
	ID         int
	UUID       string
	Token      string
	UserID     string
	LastActive string
	ExpiryAt   string
	CreatedAt  string
	UpdatedAt  string
}
