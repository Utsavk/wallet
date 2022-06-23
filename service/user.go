package service

import "wallet/models"

type User struct{}

func (u *User) GetUserDetailsByID(id string) *models.User {
	return models.FetchUserByID(id)
}
