package repository

import (
	"fmt"
	"testing"
	"wallet/errors"
	"wallet/models"
)

func TestCreateDBUser(t *testing.T) {
	type args struct {
		userInput *CreateUserInputData
	}
	type testData struct {
		name        string
		args        args
		expectedRes int64
		expectedErr *errors.Err
		wantErr     bool
	}

	var testUserData = &models.User{
		Firstname: "testFirstName",
		Lastname:  "testLastName",
		Username:  "testUserName",
		Password:  "testPassword",
		IsActive:  true,
	}
	testRole := "STAFF"
	testUserData.Role = &testRole

	data := []testData{
		{
			name: "happy flow - create user successfuly",
			args: args{
				userInput: &CreateUserInputData{
					Firstname: testUserData.Firstname,
					Lastname:  testUserData.Lastname,
					Username:  testUserData.Username,
					Password:  testUserData.Password,
					IsActive:  testUserData.IsActive,
					Role:      testUserData.Role,
				},
			},
			expectedRes: int64(testUserData.ID),
			expectedErr: nil,
			wantErr:     false,
		},
	}

	for _, tt := range data {
		fmt.Print(tt)
	}
}
