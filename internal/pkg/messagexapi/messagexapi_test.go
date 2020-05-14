package messagexapi

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/types/constants"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/util/mocks"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/util/testdata"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/pkg/logger"
)

var l *logger.Logger

func setup() {

	// Create the logger
	l = logger.CreateLogger(constants.DebugLevel)

	l.Debug("Setup completed")
}

func TestAuthenticateSuccess(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetCreated,
	}

	err := ma.Authenticate("username", "password")
	assert.NoError(t, err, "Authenticated successfully")
}

func TestAuthenticateFailure(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetUnauthorized,
	}

	err := ma.Authenticate("username", "password")
	assert.Error(t, err, "Authentication failed")
	assert.Contains(t, err.Error(), "API key and/or secret are incorrect", "Unable to log user in")
}

func TestAuthenticateFailureGarbageResponse(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	mjr := testdata.GetTestMJR()
	mjr.From = nil

	err := ma.SendMail(mjr)
	assert.Error(t, err, "Mail not sent successfully")
	assert.Contains(t, err.Error(), "invalid character", "Invalid response")
}

func TestAuthenticateURLNotFound(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetNotFound,
	}

	err := ma.Authenticate("username", "password")
	assert.Error(t, err, "Not authenticated successfully")
	assert.Contains(t, err.Error(), "error code 404", "Unable to log user in")
}

func TestSendMailSuccess(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetOk,
	}

	mjr := testdata.GetTestMJR()

	err := ma.SendMail(mjr)
	assert.NoError(t, err, "Mail sent successfully")
}

func TestSendMailFailure(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetUnprocessableEntity,
	}

	mjr := testdata.GetTestMJR()
	mjr.From = nil

	err := ma.SendMail(mjr)
	assert.Error(t, err, "Mail not sent successfully")
	assert.Contains(t, err.Error(), "failed API request, error code 422", "Invalid response")
}

func TestSendMailFailureGarbageResponse(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetGarbageResponse,
	}

	mjr := testdata.GetTestMJR()
	mjr.From = nil

	err := ma.SendMail(mjr)
	assert.Error(t, err, "Mail not sent successfully")
	assert.Contains(t, err.Error(), "invalid character", "Invalid response")
}

func TestSendMailURLNotFound(t *testing.T) {
	setup()

	ma := CreateMessageXAPI(l, constants.APIHost)
	ma.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetNotFound,
	}

	mjr := testdata.GetTestMJR()

	err := ma.SendMail(mjr)
	assert.Error(t, err, "Mail not sent successfully")
	assert.Contains(t, err.Error(), "error code 404", "Unable to send mail")
}