package models

type User struct {
	ID        int
	UUID      string
	Firstname string
	Lastname  string
	Username  string
	Password  string
	IsActive  bool
	Role      *string
	CreatedAt string
	UpdatedAt *string
	CreatedBy *string
	UpdatedBy *string
}
