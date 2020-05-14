package messagex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/messagex/go-messagex/internal/pkg/messagexapi"
	"github.com/messagex/go-messagex/internal/types/api"
	"github.com/messagex/go-messagex/internal/types/constants"
	"github.com/messagex/go-messagex/internal/util/mocks"
	"github.com/messagex/go-messagex/internal/util/testdata"
	"github.com/messagex/go-messagex/pkg/logger"
)

var l *logger.Logger

func setup(t *testing.T) *APIClient {

	// Create the logger
	l = logger.CreateLogger(constants.DebugLevel)

	ma := messagexapi.CreateMessageXAPI(l, constants.APIHost)

	ac := CreateAPIClient("api_key", "api_secret")

	ac.api = ma

	l.Debug("Setup completed")

	return ac
}

func TestCreateAPIClient(t *testing.T) {
	setup(t)
}

func TestLoginSuccess(t *testing.T) {
	ac := setup(t)

	// Set the client to a mocked one
	ac.api.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetCreated,
	}

	err := ac.Login()
	assert.NoError(t, err, "No error while creating MessageX API client")
}

func TestLoginFailure(t *testing.T) {
	ac := setup(t)

	// Set the client to a mocked one
	ac.api.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetUnauthorized,
	}

	err := ac.Login()
	assert.Error(t, err, "Unable to login")
	assert.Contains(t, err.Error(), "unable to log user in error code 401", "Login failure has an expected message")
}

func TestCreateMail(t *testing.T) {
	ac := setup(t)

	em := ac.CreateEmail()
	assert.ObjectsAreEqual(em, &api.Email{})
}


func TestSendEmailSuccess(t *testing.T) {
	ac := setup(t)

	ac.api.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetOk,
	}

	mjr := testdata.GetTestMJR()

	err := ac.SendEmail(mjr)
	assert.NoError(t, err, "Mail sent successfully")
}

func TestSendMailFailure(t *testing.T) {
	ac := setup(t)

	ac.api.Client = mocks.MockHttpClient{
		DoFunc: mocks.GetUnprocessableEntity,
	}

	mjr := testdata.GetTestMJR()
	mjr.From = nil

	err := ac.SendEmail(mjr)
	assert.Error(t, err, "Mail not sent successfully")
	assert.Contains(t, err.Error(), "failed API request, error code 422", "Invalid response")
}
