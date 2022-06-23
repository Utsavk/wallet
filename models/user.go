package models

import (
	"wallet/db/mysql"
)

type User struct {
	ID       string
	WalletID string
}

func (u *User) Save() {

}

func FetchUserByID(id string) *User {
	results, err := mysql.Conn.DB.Query("select * from user where id=?", id)
	if err != nil {
		return nil
	}
	var user User

	for results.Next() {

		err = results.Scan(&user.ID, &user.WalletID)
		if err != nil {
			continue
		}
	}
	return &user
}
