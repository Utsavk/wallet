package middleware

import (
	"fmt"
	"testing"
	"wallet/context"
	"wallet/errors"
	"wallet/logs"
	"wallet/models"
	"wallet/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func getVerifyAuthMockContext(token string) *context.Ctx {
	fctx := &fasthttp.RequestCtx{
		Request: fasthttp.Request{},
	}
	fctx.Request.Header.Add("Authorization", token)
	return context.GetMockContext(fctx, nil)
}

func TestVerifyAuth(t *testing.T) {
	testToken := "abcd"
	var testUserId uint = 1

	type args struct {
		token         string
		mockSessionFn func(session *mocks.SessionRepoInterface)
	}

	type authTest struct {
		name              string
		err               *errors.Err
		args              args
		getSessionFnCount int
	}

	testData := []authTest{
		{
			name: "pass filter for valid token",
			err:  nil,
			args: args{
				token: testToken,
				mockSessionFn: func(session *mocks.SessionRepoInterface) {
					session.On("GetSessionByToken", testToken).Return(&models.Session{
						Token:    testToken,
						ExpiryAt: "2050-01-02 00:00:00",
						UserID:   testUserId,
					}, nil)
				},
			},
			getSessionFnCount: 1,
		}, {
			name: "return error for expired token",
			args: args{
				token: testToken,
				mockSessionFn: func(session *mocks.SessionRepoInterface) {
					session.On("GetSessionByToken", testToken).Return(&models.Session{
						Token:    testToken,
						ExpiryAt: "1970-01-02 00:00:00",
						UserID:   testUserId,
					}, nil)
				},
			},
			err: &errors.Err{
				LogMessage: fmt.Sprintf("auth token %s is expired", testToken),
				UserID:     testUserId,
			},
			getSessionFnCount: 1,
		}, {
			name: "return error for blank token",
			args: args{
				token: "",
				mockSessionFn: func(session *mocks.SessionRepoInterface) {
					session.On("GetSessionByToken", mock.Anything).Return(&models.Session{}, nil)
				},
			},
			err: &errors.Err{
				LogMessage: "auth token is not provided",
			},
			getSessionFnCount: 0,
		},
	}

	for _, tt := range testData {
		logs.Print(tt.name)
		mockSession := &mocks.SessionRepoInterface{}
		authMw := AuthMw{
			mockSession,
		}
		mockCtx := getVerifyAuthMockContext(tt.args.token)
		tt.args.mockSessionFn(mockSession)
		gotErr := authMw.VerifyAuth(mockCtx)
		assert.Equal(t, tt.err, gotErr)
		mockSession.AssertNumberOfCalls(t, "GetSessionByToken", tt.getSessionFnCount)
	}
}
