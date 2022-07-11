package repository

import (
	"database/sql"
	"fmt"
	"wallet/context"
	"wallet/db/mysql"
	"wallet/errors"
	"wallet/models"
	"wallet/utils"

	"github.com/google/uuid"
)

type UserRepoInterface interface {
	GetDBUserByID(ctx *context.Ctx, id int) (*models.User, *errors.Err)
	CreateDBUser(ctx *context.Ctx, userInput *CreateUserInputData) (int64, *errors.Err)
}

type CreateUserInputData struct {
	Firstname string
	Lastname  string
	Username  string
	Password  string
	IsActive  bool
	Role      *string
}

type UserRepo struct{}

func (u *UserRepo) GetDBUserByID(ctx *context.Ctx, id int) (*models.User, *errors.Err) {
	sqlQuery := fmt.Sprintf("SELECT * from %s where id=?", models.USER_TABLE)
	results, err := mysql.Conn.DB.Query(sqlQuery, id)
	if err != nil {
		return nil, errors.NewError(err, fmt.Sprintf("user with id %d could not be fetched", id), ctx.User)
	}
	var user = models.User{}

	var role, updatedAt, createdBy, updatedBy sql.NullString
	for results.Next() {

		err = results.Scan(
			&user.UUID,
			&user.Firstname,
			&user.Lastname,
			&user.Username,
			&user.Password,
			&user.IsActive,
			&role,
			&user.CreatedAt,
			&updatedAt,
			&createdBy,
			&updatedBy,
			&user.ID,
		)
		if err != nil {
			return nil, errors.NewError(err, fmt.Sprintf("user with id %d could not be fetched", id), ctx.User)
		}
		user.Role = &role.String
		user.UpdatedAt = &updatedAt.String
		user.CreatedBy = &createdBy.String
		user.CreatedBy = &updatedBy.String
	}
	return &user, nil
}

func (u *UserRepo) CreateDBUser(ctx *context.Ctx, userInput *CreateUserInputData) (int64, *errors.Err) {
	uuid, _ := uuid.NewUUID()
	user := &models.User{
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Username:  userInput.Username,
		Password:  userInput.Password,
		IsActive:  userInput.IsActive,
		Role:      userInput.Role,
	}

	user.CreatedAt = utils.ClockObj.GetCurrentTime()

	sqlQuery := `INSERT INTO ` + models.USER_TABLE + `(
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
		return -1, errors.NewError(err, fmt.Sprintf("error in creating new user %s", user.Username), ctx.User)
	}

	lastId, err := results.LastInsertId()
	if err != nil {
		return -1, errors.NewError(err, fmt.Sprintf("error in creating new user %s", user.Username), ctx.User)
	}

	return lastId, nil
}
