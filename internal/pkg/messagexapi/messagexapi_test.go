package messagexapi
//
//import (
//	"bytes"
//	"io/ioutil"
//	"messagex.com/mailserver/internal/util/testdata"
//	"net/http"
//	"os"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	"messagex.com/mailserver/internal/util/mocks"
//	"messagex.com/mailserver/pkg/env"
//	"messagex.com/mailserver/pkg/logger"
//)
//
//var e *env.Env
//var l *logger.Logger
//
//func setup(t *testing.T) {
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
//	l.Debug("Setup completed")
//}
//
//func TestCreateDBConSuccess(t *testing.T) {
//	setup(t)
//
//	ma, err := CreateMessageXAPI(e, l, "localhost")
//	if noError := assert.NoError(t, err, "No error while creating the DB"); noError {
//		assert.NotNil(t, ma, "MessageX API has been created")
//	}
//}
//
//func TestAuthenticateSuccess(t *testing.T) {
//	setup(t)
//
//	ma, err := CreateMessageXAPI(e, l, "localhost")
//	if noError := assert.NoError(t, err, "No error while creating the DB"); noError {
//		assert.NotNil(t, ma, "MessageX API has been created")
//	}
//
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
//	err = ma.Authenticate("username", "password")
//	assert.NoError(t, err, "Authenticated successfully")
//}
//
//func TestAuthenticateFailure(t *testing.T) {
//	setup(t)
//
//	ma, err := CreateMessageXAPI(e, l, "localhost")
//	if noError := assert.NoError(t, err, "No error while creating the DB"); noError {
//		assert.NotNil(t, ma, "MessageX API has been created")
//	}
//
//	json := `
//{
//	"message": "The given data was invalid",
//	"errors": {
//	"apiKey/apiSecret": "API key and/or secret are incorrect."
//}
//`
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
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
//	err = ma.Authenticate("username", "password")
//	assert.Error(t, err, "Authentication failed")
//	assert.Contains(t, err.Error(), "API key and/or secret are incorrect", "Unable to log user in")
//}
//
//func TestSendMailSuccess(t *testing.T) {
//	setup(t)
//
//	ma, err := CreateMessageXAPI(e, l, "localhost")
//	if noError := assert.NoError(t, err, "No error while creating the DB"); noError {
//		assert.NotNil(t, ma, "MessageX API has been created")
//	}
//
//	mjr := testdata.GetTestMJR()
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
//	emailEnabled := e.Get("APP_EMAIL_ENABLED")
//	_ = os.Setenv("APP_EMAIL_ENABLED", "1")
//
//	err = ma.SendMail(mjr)
//	assert.NoError(t, err, "Mail sent successfully")
//
//	_ = os.Setenv("APP_EMAIL_ENABLED", emailEnabled)
//}
//
//func TestSendMailFailure(t *testing.T) {
//	setup(t)
//
//	ma, err := CreateMessageXAPI(e, l, "localhost")
//	if noError := assert.NoError(t, err, "No error while creating the DB"); noError {
//		assert.NotNil(t, ma, "MessageX API has been created")
//	}
//
//	mjr := testdata.GetTestMJR()
//	json := `{"success": false}`
//
//	doFunc := func(*http.Request) (*http.Response, error) {
//		// create a new reader with that JSON
//		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
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
//	err = ma.SendMail(mjr)
//	assert.Error(t, err, "Mail not sent successfully")
//
//	_ = os.Setenv("APP_EMAIL_ENABLED", emailEnabled)
//}