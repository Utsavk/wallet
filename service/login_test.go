package service

import (
	"testing"
	"wallet/errors"
	"wallet/models"
	"wallet/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	testUser := models.User{}
	testSession := models.Session{
		UserID: testUser.ID,
		Token:  "authHeaderValue",
	}
	type args struct {
		username string
		password string
	}
	type testData struct {
		name                   string
		args                   args
		mockFn                 func() *mocks.SessionRepoInterface
		wantErr                bool
		expectedMsg            string
		expectedAuthHeader     string
		err                    *errors.Err
		getUserCallCount       int
		createSessionCallCount int
	}

	data := []testData{
		{
			name: "happy flow - successful login",
			args: args{
				username: testUser.Username,
				password: testUser.Password,
			},
			mockFn: func() *mocks.SessionRepoInterface {
				mockSessionRepo := mocks.SessionRepoInterface{}
				mockSessionRepo.On("CreateDBSession", testUser.ID).Return(&testSession)
				return &mockSessionRepo // return userRepoInterface also to call GetUserByID
			},
			wantErr:                false,
			expectedMsg:            "logged in successfuly",
			expectedAuthHeader:     testSession.Token,
			err:                    nil,
			getUserCallCount:       1,
			createSessionCallCount: 1,
		},
	}

	for _, tt := range data {
		var loginService = &LoginService{}
		sessionRepoMock := tt.mockFn()
		loginService.sessionRepo = sessionRepoMock
		actualMsg, actualHeader, actualErr := loginService.Login(LoginRequest{
			Username: tt.args.username,
			Password: tt.args.password,
		})
		assert.EqualValues(t, tt.err, actualErr)
		if !tt.wantErr {
			assert.Equal(t, tt.expectedMsg, actualMsg)
			assert.Equal(t, tt.expectedAuthHeader, actualHeader)
		} else {
			assert.Nil(t, actualErr)
		}
		assert.True(t, sessionRepoMock.AssertNumberOfCalls(t, "CreateDBSession", tt.createSessionCallCount))
	}
}
