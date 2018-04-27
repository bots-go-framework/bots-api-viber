package viberbotapi

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/strongo/bots-api-viber/viberinterface"
	"io/ioutil"
	"net/http"
)

const (
	// ViberAPIBaseURL is a base API url
	ViberAPIBaseURL = "https://chatapi.viber.com/pa/"
)

// ViberBotAPI allows you to interact with the Viber Bot API.
type ViberBotAPI struct {
	httpClient *http.Client
	token      string
}

// NewViberBotAPIWithHTTPClient creates new API provider with HTTP client
func NewViberBotAPIWithHTTPClient(token string, httpClient *http.Client) *ViberBotAPI {
	return &ViberBotAPI{token: token, httpClient: httpClient}
}

// SetWebhook sets webhook to the specified URL
func (botApi ViberBotAPI) SetWebhook(url string, eventTypes []string) (response viberinterface.SetWebhookResponse, err error) {
	var responseBody []byte
	if _, responseBody, err = botApi.send(
		&viberinterface.SetWebhookMessage{
			ViberAuth:  viberinterface.ViberAuth{Token: botApi.token},
			Url:        url,
			EventTypes: eventTypes,
		},
	); err != nil {
		return
	}

	if err = ffjson.UnmarshalFast(responseBody, &response); err != nil {
		err = errors.Wrap(err, "Failed to unmarshal response body to JSON")
		return
	}
	return
}

// SendMessage sends messages to Viber
func (botApi ViberBotAPI) SendMessage(m viberinterface.MessageToReceiver) (requestBody []byte, response viberinterface.SendMessageResponse, err error) {
	m.SetType(m.GetType())
	m.SetToken(botApi.token)
	var responseBody []byte
	if requestBody, responseBody, err = botApi.send(m); err != nil {
		return
	}

	if err = ffjson.UnmarshalFast(responseBody, &response); err != nil {
		err = errors.Wrap(err, "Failed to unmarshal response body to JSON")
	}
	return
}

func (botApi ViberBotAPI) send(m viberinterface.MessageToViberEndpoint) (requestBody []byte, responseBody []byte, err error) {
	if requestBody, err = ffjson.MarshalFast(m); err != nil {
		ffjson.Pool(requestBody)
		return
	}
	var resp *http.Response
	endpointURL := ViberAPIBaseURL + m.Endpoint()
	resp, err = botApi.httpClient.Post(endpointURL, "applicaiton/json", bytes.NewReader(requestBody))
	ffjson.Pool(requestBody)
	if err != nil {
		return
	}
	if responseBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}
