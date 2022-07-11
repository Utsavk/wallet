package service

import (
	"wallet/context"
	"wallet/errors"
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
	ID        uint
	UUID      string
	Firstname string
	Lastname  string
	Username  string
	IsActive  bool
	Role      *string
}

type UserServiceInterface interface {
	repository.UserRepoInterface
	GetUserDetailsByID(ctx *context.Ctx, id string) (*User, *errors.Err)
	CreateUser(ctx *context.Ctx, args NewUserArgs) (int64, *errors.Err)
}

type UserService struct {
	UserRepo repository.UserRepoInterface
}

func (u *UserService) GetUserDetailsByID(ctx *context.Ctx, id int) (*User, *errors.Err) {
	dbUser, err := u.UserRepo.GetDBUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        dbUser.ID,
		UUID:      dbUser.UUID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		IsActive:  dbUser.IsActive,
		Role:      dbUser.Role,
	}, nil
}

func (u *UserService) CreateUser(ctx *context.Ctx, args NewUserArgs) (int64, *errors.Err) {
	newUserId, err := u.UserRepo.CreateDBUser(ctx, &repository.CreateUserInputData{
		Firstname: args.Firstname,
		Lastname:  args.Lastname,
		Username:  args.Username,
		Password:  args.Password,
		IsActive:  args.IsActive,
		Role:      &args.Role,
	})
	if err != nil {
		return -1, err
	}
	return newUserId, nil
}
