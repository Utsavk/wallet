package service

import (
	"testing"
	"wallet/errors"
	"wallet/models"
	"wallet/repository/mocks"
	"wallet/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	type args struct {
		userID uint
		mockFn func(userID uint, session *models.Session, err *errors.Err) *mocks.SessionRepoInterface
	}

	type testData struct {
		name string
		args args
		err  *errors.Err
		resp *Session
	}

	var testUserId uint = 1
	data := []testData{
		{
			name: "happy flow - create session object",
			args: args{
				userID: testUserId,
				mockFn: func(userID uint, session *models.Session, err *errors.Err) *mocks.SessionRepoInterface {
					sessionRepoMock := &mocks.SessionRepoInterface{}
					sessionRepoMock.On("CreateDBSession", testUserId).Return(session, err)
					return sessionRepoMock
				},
			},
			err: nil,
			resp: &Session{
				ID:     1,
				UUID:   "testuuid",
				UserID: testUserId,
				Token:  "testtoken",
			},
		},
	}

	for _, tt := range data {
		sessionService := &SessionService{}
		sessionService.sessionRepo = tt.args.mockFn(testUserId, &models.Session{
			ID:        1,
			UUID:      "testuuid",
			UserID:    testUserId,
			Token:     "testtoken",
			CreatedAt: utils.GetCurrentTime(),
		}, nil)
		actualSession, actualErr := sessionService.CreateSession(tt.args.userID)
		assert.Nil(t, actualErr)
		assert.Equal(t, tt.resp.Token, actualSession.Token)
		assert.Equal(t, tt.resp.ID, actualSession.ID)
		assert.Equal(t, tt.resp.UUID, actualSession.UUID)
		assert.Equal(t, tt.resp.UserID, actualSession.UserID)
	}
}
