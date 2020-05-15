package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// MockHttpClient is the mock client
type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	//return GetDoFunc(req)
	return m.DoFunc(req)
}

func GetOk (*http.Request) (*http.Response, error) {
	// create a new reader with that JSON
	json := `{"success": true}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}


func GetUnauthorized(*http.Request) (*http.Response, error) {

	json := `
{
	"message": "The given data was invalid",
	"errors": {
	"apiKey/apiSecret": "API key and/or secret are incorrect."
}
`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       r,
	}, nil
}


func GetCreated(*http.Request) (*http.Response, error) {

	json := `
{
	"data": {
		"bearerToken":"9UEvJsGKagOQ3JHSpTw3e727Cnrle4rnqkYO4jorMR3t5TYaeUnYPYNe90Cy9BA7N2b2rX2KCsTuRxrfQ2Td7hgOtKll1Sl1WfofBBYnP5Y5RyoVZsT6A2s3Ef0kPYyK",
		"refreshToken":"p7lh5nhVoNKjP8vKJzTIoEJo9aopIT0DIswEq2RdTgkriEc0LDNAV2lnsq1MV3MvAUMZJoHmGuBbEzCmurJBznHHk7Dm1YkY86U3TCHIQaJoxRBMwr9D8gkPFu2Gjq9H",
		"expiresAt":"2021-02-21 03:22:36",
		"apiKeyId":1,
		"updatedAt":"2020-02-21 03:22:36",
		"createdAt":"2020-02-21 03:22:36",
		"id":55,
		"createdAtEpoch": 1582255356,
		"updatedAtEpoch": 1582255356
	}
}
`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusCreated,
		Body:       r,
	}, nil
}

func GetUnprocessableEntity(*http.Request) (*http.Response, error) {
	json := `
{
	"message": "The given data was invalid.",
	"errors": {
		"from": [
			"The from field is required."
		]
	}
}
`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusUnprocessableEntity,
		Body:       r,
	}, nil
}


func GetGarbageResponse(*http.Request) (*http.Response, error) {
	json := `{garbage`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}

func GetNotFound(*http.Request) (*http.Response, error) {

	json := `
{
	"message": ""
}
`

	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       r,
	}, nil
}