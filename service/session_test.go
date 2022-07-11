package service

import (
	"testing"
	"wallet/context"
	"wallet/errors"
	"wallet/models"
	repoMocks "wallet/repository/mocks"
	"wallet/utils"
	utilsMocks "wallet/utils/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSession(t *testing.T) {
	mockedCurrentTimeStr := "2021-07-10 10:00:07"

	type args struct {
		userID uint
		mockFn func() *repoMocks.SessionRepoInterface
	}

	type testData struct {
		name                  string
		args                  args
		err                   *errors.Err
		expected              *Session
		createSessionNumCalls int
		getSessionNumCalls    int
		wantErr               bool
	}

	var testUserId uint = 1

	genuineClockObj := setUpCurrentTimeMock(mockedCurrentTimeStr)
	dummyCurrentTime := utils.ClockObj.GetCurrentTime()

	data := []testData{
		{
			name: "happy flow - create session object",
			args: args{
				userID: testUserId,
				mockFn: func() *repoMocks.SessionRepoInterface {
					sessionRepoMock := &repoMocks.SessionRepoInterface{}
					sessionRepoMock.On("CreateDBSession", mock.Anything, testUserId).Return(int64(testUserId), nil)
					sessionRepoMock.On("GetDBSessionByUserID", mock.Anything, testUserId).Return(&models.Session{
						ID:        1,
						UUID:      "testuuid",
						UserID:    testUserId,
						Token:     "testtoken",
						CreatedAt: dummyCurrentTime,
						ExpiryAt:  dummyCurrentTime,
					}, nil)
					return sessionRepoMock
				},
			},
			err: nil,
			expected: &Session{
				ID:       1,
				UUID:     "testuuid",
				UserID:   testUserId,
				Token:    "testtoken",
				ExpiryAt: dummyCurrentTime,
			},
			createSessionNumCalls: 1,
			getSessionNumCalls:    1,
			wantErr:               false,
		},
		{
			name: "throws error on invalid user id",
			args: args{
				userID: 0,
				mockFn: func() *repoMocks.SessionRepoInterface {
					return nil
				},
			},
			err:                   errors.NewError(nil, "invalid userid", nil),
			expected:              nil,
			createSessionNumCalls: 0,
			getSessionNumCalls:    0,
			wantErr:               true,
		},
	}

	mockContext := context.GetMockContext(nil, nil)
	for _, tt := range data {
		sessionService := &SessionService{}
		sessionRepoMock := tt.args.mockFn()
		sessionService.sessionRepo = sessionRepoMock
		actualSession, actualErr := sessionService.CreateSession(mockContext, tt.args.userID)
		assert.EqualValues(t, tt.err, actualErr)
		if !tt.wantErr {
			assert.Equal(t, tt.expected.Token, actualSession.Token)
			assert.Equal(t, tt.expected.ID, actualSession.ID)
			assert.Equal(t, tt.expected.UUID, actualSession.UUID)
			assert.Equal(t, tt.expected.UserID, actualSession.UserID)
			assert.Equal(t, tt.expected.ExpiryAt, actualSession.ExpiryAt)
			assert.True(t, sessionRepoMock.AssertNumberOfCalls(t, "CreateDBSession", tt.createSessionNumCalls))
			assert.True(t, sessionRepoMock.AssertNumberOfCalls(t, "GetDBSessionByUserID", tt.getSessionNumCalls))
		} else {
			assert.NotNil(t, actualErr)
			assert.EqualValues(t, tt.err, actualErr)
		}
	}
	utils.ClockObj = genuineClockObj
}

func setUpCurrentTimeMock(mockedCurrentTimeStr string) utils.ClockInterface {
	genuineClockObj := utils.ClockObj
	clockObjMock := &utilsMocks.ClockInterface{}
	clockObjMock.On("GetCurrentTime").Return(mockedCurrentTimeStr)
	utils.ClockObj = clockObjMock
	return genuineClockObj
}
