package messagex
//
//import (
//	"bytes"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"strings"
//	"testing"
//
//	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/pkg/smtpserver"
//
//	"github.com/stretchr/testify/assert"
//
//	"messagex.com/mailserver/internal/pkg/emailparser"
//	"messagex.com/mailserver/internal/pkg/messagexapi"
//	"messagex.com/mailserver/internal/util/mocks"
//	"messagex.com/mailserver/pkg/env"
//	"messagex.com/mailserver/pkg/logger"
//)
//
//var e *env.Env
//var l *logger.Logger
//
//func setup(t *testing.T) *messagexapi.MessageXAPI {
//
//	var err error
//
//	// Create the environment
//	e, err = env.CreateEnvFromFile("../../../.env")
//	if err != nil {
//		t.Fatalf("Unable to create environment: %v", err)
//	}
//
//	// Create the logger
//	l = logger.CreateLogger(e.Get("APP_DEBUG_LEVEL"))
//
//	ma, err := messagexapi.CreateMessageXAPI(e, l, "")
//	assert.NoError(t, err, "No error while creating the DB")
//
//	l.Debug("Setup completed")
//
//	return ma
//}
//
//func TestCreateSMTPServerSuccess(t *testing.T) {
//	ma := setup(t)
//
//	ep := emailparser.CreateEmailParser(e, l)
//
//	sr, err := smtpserver.CreateSMTPServer(e, l, ma, ep)
//	if noError := assert.NoError(t, err, "No error while creating the SMTP server"); noError {
//		assert.NotNil(t, sr, "SMTP server has been created")
//	}
//}
//
//
//func GetSMTPSession(t *testing.T, ma *messagexapi.MessageXAPI) (smtp.Session, error) {
//	json := `
//{
//	"data": {
//		"bearerToken":"9UEvJsGKagOQ3JHSpTw3e727Cnrle4rnqkYO4jorMR3t5TYaeUnYPYNe90Cy9BA7N2b2rX2KCsTuRxrfQ2Td7hgOtKll1Sl1WfofBBYnP5Y5RyoVZsT6A2s3Ef0kPYyK",
//		"refreshToken":"p7lh5nhVoNKjP8vKJzTIoEJo9aopIT0DIswEq2RdTgkriEc0LDNAV2lnsq1MV3MvAUMZJoHmGuBbEzCmurJBznHHk7Dm1YkY86U3TCHIQaJoxRBMwr9D8gkPFu2Gjq9H",
//		"expiresAt":"2021-02-21 03:22:36",
//		"apiKeyId":1,
//		"updatedAt":"2020-02-21 03:22:36",
//		"createdAt":"2020-02-21 03:22:36",
//		"id":55,
//		"createdAtEpoch": 1582255356,
//		"updatedAtEpoch": 1582255356
//	}
//}
//`
//
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
//		return &http.Response{
//			StatusCode: http.StatusCreated,
//			Body:       r,
//		}, nil
//	}
//
//	ma.Client = mocks.MockHttpClient{
//		DoFunc: doFunc,
//	}
//
//	ep := emailparser.CreateEmailParser(e, l)
//
//	sr, err := smtpserver.CreateSMTPServer(e, l, ma, ep)
//	if noError := assert.NoError(t, err, "No error while creating the SMTP server"); noError {
//		assert.NotNil(t, sr, "SMTP server has been created")
//	}
//
//	s, err := sr.Backend.Login(nil, "user", "pass")
//	if noError := assert.NoError(t, err, "No error while logging in to the SMTP server"); noError {
//		assert.NotNil(t, s, "SMTP server has been created")
//	}
//
//	return s, nil
//}
//
//func TestLoginSuccess(t *testing.T) {
//	ma := setup(t)
//
//	_, _ = GetSMTPSession(t, ma)
//}
//
//func TestLoginFailure(t *testing.T) {
//	ma := setup(t)
//
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(bytes.NewReader([]byte("Unauthorized.")))
//		return &http.Response{
//			StatusCode: http.StatusUnauthorized,
//			Body:       r,
//		}, nil
//	}
//
//	ma.Client = mocks.MockHttpClient{
//		DoFunc: doFunc,
//	}
//
//	ep := emailparser.CreateEmailParser(e, l)
//
//	sr, err := smtpserver.CreateSMTPServer(e, l, ma, ep)
//	if noError := assert.NoError(t, err, "No error while creating the SMTP server"); noError {
//		assert.NotNil(t, sr, "SMTP server has been created")
//	}
//
//	s, err := sr.Backend.Login(nil, "user", "pass")
//	assert.Error(t, err, "Unable to log in to the SMTP server")
//	assert.Nil(t, s, "No SMTP session was created")
//}
//
//func TestAnonymousLogin(t *testing.T) {
//	ma := setup(t)
//
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(bytes.NewReader([]byte("Unauthorized.")))
//		return &http.Response{
//			StatusCode: http.StatusUnauthorized,
//			Body:       r,
//		}, nil
//	}
//
//	ma.Client = mocks.MockHttpClient{
//		DoFunc: doFunc,
//	}
//
//	ep := emailparser.CreateEmailParser(e, l)
//
//	sr, err := smtpserver.CreateSMTPServer(e, l, ma, ep)
//	if noError := assert.NoError(t, err, "No error while creating the SMTP server"); noError {
//		assert.NotNil(t, sr, "SMTP server has been created")
//	}
//
//	s, err := sr.Backend.AnonymousLogin(nil)
//	assert.Error(t, err, "Unable to log in to the SMTP server")
//	assert.Nil(t, s, "No SMTP session was created")
//}
//
//
//func TestMail(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//
//	err := s.Mail("from@example.com", smtp.MailOptions{})
//	assert.NoError(t, err, "No error from the Mail function")
//}
//
//func TestRcpt(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//
//	err := s.Rcpt("from@example.com")
//	assert.NoError(t, err, "No error from the Rcpt function")
//}
//
//func TestDataSuccess(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//
//	json := `{"success": true}`
//
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
//		return &http.Response{
//			StatusCode: http.StatusOK,
//			Body:       r,
//		}, nil
//	}
//
//	ma.Client = mocks.MockHttpClient{
//		DoFunc: doFunc,
//	}
//
//	err := s.Data(testdata.GetTestEmailReader())
//	assert.NoError(t, err, "No error from the Mail function")
//}
//
//func TestDataUnauthorized(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(strings.NewReader("Unauthorized."))
//		return &http.Response{
//			StatusCode: http.StatusUnauthorized,
//			Body:       r,
//		}, nil
//	}
//
//	ma.Client = mocks.MockHttpClient{
//		DoFunc: doFunc,
//	}
//
//	emailEnabled := e.Get("APP_EMAIL_ENABLED")
//	_ = os.Setenv("APP_EMAIL_ENABLED", "1")
//
//	err := s.Data(testdata.GetTestEmailReader())
//	assert.Error(t, err, "Unauthorized error from the Mail function")
//
//	_ = os.Setenv("APP_EMAIL_ENABLED", emailEnabled)
//}
//
//func TestDataUnparsableEmail(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//
//	err := s.Data(strings.NewReader("Unparseable Data"))
//	assert.Error(t, err, "No error from the Mail function")
//	assert.Contains(t, err.Error(), "Failed to ReadParts", "The error is as expected")
//}
//
//func TestReset(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//	s.Reset()
//}
//
//func TestLogout(t *testing.T) {
//	ma := setup(t)
//
//	s, _ := GetSMTPSession(t, ma)
//
//	err := s.Logout()
//	assert.NoError(t, err, "No error from the Logout function")
//}