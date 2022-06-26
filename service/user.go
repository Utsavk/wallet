package service

import (
	"wallet/models"
	"wallet/repository"
)

type NewUserArgs struct {
	Firstname string
	Lastname  string
	Username  string
	Password  string
	IsActive  bool
	Role      string
}

type User struct {
	ID        int
	UUID      string
	Firstname string
	Lastname  string
	Username  string
	Password  string
	IsActive  bool
	Role      string
}

type UserServiceInterface interface {
	GetUserDetailsByID(id string) *models.User
	CreateUser(args NewUserArgs)
}

type UserService struct{}

func (u *UserService) GetUserDetailsByID(id int) *User {
	dbUser := repository.GetUserByID(id)
	return &User{
		ID:        dbUser.ID,
		UUID:      dbUser.UUID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
	}
}

func (u *UserService) CreateUser(args NewUserArgs) *User {
	dbUser := repository.CreateUser(&models.User{
		Firstname: args.Firstname,
		Lastname:  args.Lastname,
		Username:  args.Username,
		Password:  args.Password,
		IsActive:  args.IsActive,
		Role:      &args.Role,
	})
	if dbUser == nil {
		return nil
	}
	return &User{
		ID:        dbUser.ID,
		UUID:      dbUser.UUID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
	}
}
