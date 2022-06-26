package middleware

import (
	"fmt"
	"testing"
	"wallet/errors"
	"wallet/logs"
	"wallet/models"
	"wallet/repository/mocks"
	"wallet/wcontext"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func getVerifyAuthMockContext(token string) *wcontext.Context {
	fctx := &fasthttp.RequestCtx{
		Request: fasthttp.Request{},
	}
	fctx.Request.Header.Add("Authorization", token)
	return &wcontext.Context{
		Fctx: fctx,
	}
}

func TestVerifyAuth(t *testing.T) {
	testToken := "abcd"

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
						ExpiryAt: "",
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
						ExpiryAt: "",
					}, nil)
				},
			},
			err: &errors.Err{
				LogMessage: fmt.Sprintf("auth token %s is expired", testToken),
			},
			getSessionFnCount: 1,
		}, {
			name: "return error for blank token",
			args: args{
				token: "",
				mockSessionFn: func(session *mocks.SessionRepoInterface) {
					session.On("GetSessionByToken", mock.Anything).Return(&models.Session{
						Token:    testToken,
						ExpiryAt: "",
					}, nil)
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
		gotErr := authMw.VerifyAuth(mockCtx.Fctx)
		assert.Equal(t, tt.err, gotErr)
		mockSession.AssertNumberOfCalls(t, "GetSessionByToken", tt.getSessionFnCount)
	}
}
