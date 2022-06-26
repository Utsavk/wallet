package repository

import (
	"database/sql"
	"fmt"
	"time"
	"wallet/db/mysql"
	"wallet/logs"
	"wallet/models"

	"github.com/google/uuid"
)

const USER_TABLE = "user"

func GetUserByID(id int) *models.User {
	sql := fmt.Sprintf("SELECT * from %s where id=?", USER_TABLE)
	results, err := mysql.Conn.DB.Query(sql, id)
	if err != nil {
		logs.Print(err.Error())
		return nil
	}
	var user models.User

	for results.Next() {

		err = results.Scan(&user.ID)
		if err != nil {
			logs.Print(err.Error())
			continue
		}
	}
	return &user
}

func CreateUser(user *models.User) *models.User {
	uuid, _ := uuid.NewUUID()
	user.CreatedAt = time.Now().Format("2006/01/02 15:04:05")

	sqlQuery := `INSERT INTO ` + USER_TABLE + `(
		uuid, 
		firstname, 
		lastname,
		username,
		password,
		isactive,
		role,
		createdat,
		updatedat,
		createdby,
		updatedby
		) 
		VALUES 
		(
			?,?,?,?,?,?,?,?,?,?,?
		)`

	var createdBy interface{}
	if user.CreatedBy == nil {
		createdBy = sql.NullString{}
	}

	var role interface{}
	if user.Role == nil {
		role = sql.NullString{}
	}

	results, err := mysql.Conn.DB.Exec(sqlQuery,
		uuid.String(),
		user.Firstname,
		user.Lastname,
		user.Username,
		user.Password,
		user.IsActive,
		role,
		user.CreatedAt,
		sql.NullString{},
		createdBy,
		sql.NullString{},
	)

	if err != nil {
		logs.Print(err.Error())
		return nil
	}

	lastId, err := results.LastInsertId()
	if err != nil {
		logs.Print(err.Error())
		return nil
	}

	return GetUserByID(int(lastId))
}