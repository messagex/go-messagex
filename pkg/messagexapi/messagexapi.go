package messagexapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/messagex/go-messagex/interface/apiclient"
	"github.com/messagex/go-messagex/pkg/logger"
	"github.com/messagex/go-messagex/types/api"
	"github.com/messagex/go-messagex/types/constants"
)

type MessageXAPI struct {
	l *logger.Logger

	APIHost     string
	BearerToken string

	Client apiclient.HTTPClient
}

// CreateMessageXAPI - create message x api layer instance*/
func CreateMessageXAPI(l *logger.Logger, apiHost string) *MessageXAPI {
	log := l.Lgr.With().Str("MessageX API Layer", "CreateMessageXAPI").Logger()

	log.Info().Msg("Creating MessageX API Instance")

	return &MessageXAPI{
		l:       l,
		APIHost: apiHost,
		Client:  &http.Client{},
	}
}

// Login handles a login command with username and password.
func (ma *MessageXAPI) Authenticate(username, password string) error {
	l := ma.l.Lgr.With().Str("MessageX API Layer", "Authenticate").Logger()

	ar := &api.AuthenticateRequest{
		APIKey:    username,
		APISecret: password,
	}

	arj, _ := json.Marshal(ar)

	apiUrl := fmt.Sprintf("%s/api/%s", ma.APIHost, constants.EndpointAuthorise)

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(arj))
	if err != nil {
		return err
	}

	// add headers
	req.Header.Add("Content-Type", constants.ContentTypeJson)
	req.Header.Add("Accept", constants.ContentTypeJson)

	// Send req using http Client
	resp, err := ma.Client.Do(req)
	if err != nil {
		l.Error().Msgf("Error while making request to %s API: %s", constants.EndpointMailSend, err.Error())
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		l.Error().Msgf("Unable to log user in error code %d, error %s", resp.StatusCode, string(respBody))
		return fmt.Errorf("unable to log user in error code %d, error %s", resp.StatusCode, string(respBody))
	}

	var auResp api.AuthenticateResponse
	err = json.Unmarshal(respBody, &auResp)
	if err != nil {
		return err
	}

	ma.BearerToken = auResp.Data.BearerToken

	l.Debug().Msgf("Logged in with username: %s", username)

	return nil
}

func (ma *MessageXAPI) SendMail(mjr *api.Email) error {
	l := ma.l.Lgr.With().Str("MessageX API Layer", "SendMail").Logger()

	mjrJson, _ := json.Marshal(mjr)

	l.Debug().Msgf(string(mjrJson))

	apiUrl := fmt.Sprintf("%s/api/%s", ma.APIHost, constants.EndpointMailSend)

	// Create a new request using http
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(mjrJson))
	if err != nil {
		l.Error().Msgf("Error creating request to mail/send API: %s", err.Error())
		return err
	}

	// add authorization header to the req
	req.Header.Add("Content-Type", constants.ContentTypeJson)
	req.Header.Add("Accept", constants.ContentTypeJson)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ma.BearerToken))

	// Send req using http Client
	resp, err := ma.Client.Do(req)
	if err != nil {
		l.Error().Msgf("Error while making request to mail/send API: %s", err.Error())
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Error().Msgf("Error while reading response from mail/send API: %s", err.Error())
		return err
	}

	if resp.StatusCode != http.StatusOK {
		l.Error().Msgf("Failed API request, error code %d, error %s", resp.StatusCode, string(respBody))
		return fmt.Errorf("failed API request, error code %d, error %s", resp.StatusCode, string(respBody))
	}

	l.Debug().Msgf("Sent an email to the API layer")

	var msResp api.MailSendResponse
	err = json.Unmarshal(respBody, &msResp)
	if err != nil {
		l.Error().Msg(err.Error())
		return err
	}

	return nil
}
