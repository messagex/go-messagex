package messagex

import (
	"fmt"
	"os"

	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/pkg/messagexapi"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/types/api"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/internal/types/constants"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/pkg/env"
	"vines.smsglobal.local/messagex/v2/sdk/go-messagex/pkg/logger"
)

type APIClient struct {
	l *logger.Logger
	e *env.Env

	apiKey, apiSecret string

	api *messagexapi.MessageXAPI
}

func CreateAPIClient(user, pass string) (*APIClient, error) {

	// Create the logger
	l := logger.CreateLogger(constants.DebugLevel)

	mxApi := &APIClient{
		l:         l,
		apiKey:    user,
		apiSecret: pass,
	}

	mxApi.apiKey = user
	mxApi.apiSecret = pass

	a, err := messagexapi.CreateMessageXAPI(l, constants.APIHost)
	if err != nil {
		fmt.Printf("Error checking env variables: %s", err.Error())
		os.Exit(1)
	}

	mxApi.api = a

	err = mxApi.Login()
	if err != nil {
		return nil, fmt.Errorf("unable to log in to the MessageX api server: %s", err.Error())
	}

	return mxApi, nil
}

// Login handles a login command with username and password.
func (c *APIClient) Login() error {
	lg := c.l.Lgr.With().Str("Mail Server", "Login").Logger()

	err := c.api.Authenticate(c.apiKey, c.apiSecret)
	if err != nil {
		return err
	}

	lg.Info().Msg("Successfully authenticated with the API layer")

	return nil
}

func (c *APIClient) CreateEmail() *api.Email {
	lg := c.l.Lgr.With().Str("Mail Server", "CreateEmail").Logger()

	lg.Debug().Msg("Creating an email message")

	return &api.Email{}
}

func (c *APIClient) SendEmail(mjr *api.Email) error {
	lg := c.l.Lgr.With().Str("MessageX", "SendEmail").Logger()

	lg.Info().Msg("Sending email to MessageX")
	lg.Debug().Msgf("%v", mjr)

	err := c.api.SendMail(mjr)
	if err != nil {
		return err
	}

	lg.Info().Msg("Done sending email")

	return nil
}