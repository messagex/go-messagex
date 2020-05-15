// Package MessageX provides the ADK to the MessageX API server
//
// Copyright 2020 SMSGlobal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package messagex

import (
	"github.com/messagex/go-messagex/pkg/logger"
	"github.com/messagex/go-messagex/pkg/messagexapi"
	"github.com/messagex/go-messagex/types/api"
	"github.com/messagex/go-messagex/types/constants"
)

type APIClient struct {
	l *logger.Logger
	apiKey, apiSecret string
	api *messagexapi.MessageXAPI
}

// CreateAPIClient - Initializes the MessageX api client library and provides a handle that can later be used to login
// to the MessageX API server and execute requests against it.
//
// Parameters:
//
//  user - the login username/api key
//  pass - the password/api secret
//
// Returns an instance of the APIClient
func CreateAPIClient(user, pass string) *APIClient {

	// Create the logger
	l := logger.CreateLogger(constants.DebugLevel)
	lg := l.Lgr.With().Str("Mail Server", "Login").Logger()

	mxApi := &APIClient{
		l:         l,
		apiKey:    user,
		apiSecret: pass,
	}

	mxApi.apiKey = user
	mxApi.apiSecret = pass

	mxApi.api = messagexapi.CreateMessageXAPI(l, constants.APIHost)

	lg.Info().Msgf("API Client created successfully")

	return mxApi
}

// Login - Logs the user in to the MessageX API server with the username and password provided in the call to CreateAPIClient.
//
// Returns an error if it occurs
func (c *APIClient) Login() error {
	lg := c.l.Lgr.With().Str("Mail Server", "Login").Logger()

	err := c.api.Authenticate(c.apiKey, c.apiSecret)
	if err != nil {
		return err
	}

	lg.Info().Msg("Successfully authenticated with the API layer")

	return nil
}

// CreateEmail - Creates an empty email object that is to be populated and sent to the MessageX API server
func (c *APIClient) CreateEmail() *api.Email {
	lg := c.l.Lgr.With().Str("Mail Server", "CreateEmail").Logger()

	lg.Debug().Msg("Creating an email message")

	return &api.Email{}
}

// SendEmail - Send the email to the MessageX API server
//
// Parameters:
//	msg - An instance on the api.Email struct that represents the email message to be sent
//
// Returns an error if it occurs
func (c *APIClient) SendEmail(msg *api.Email) error {
	lg := c.l.Lgr.With().Str("MessageX", "SendEmail").Logger()

	lg.Info().Msg("Sending email to MessageX")
	lg.Debug().Msgf("%v", msg)

	err := c.api.SendMail(msg)
	if err != nil {
		return err
	}

	lg.Info().Msg("Done sending email")

	return nil
}