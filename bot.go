package viberbotapi

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/pkg/errors"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/strongo/bots-api-viber/viberinterface"
)

const (
	VIBER_API_BASE_URL = "https://chatapi.viber.com/pa/"
)

// BotAPI allows you to interact with the Viber Bot API.
type ViberBotApi struct {
	httpClient *http.Client
	token      string
}

func NewViberBotApiWithHttpClient(token string, httpClient *http.Client) *ViberBotApi {
	return &ViberBotApi{token: token, httpClient: httpClient}
}

func (botApi ViberBotApi) SetWebhook(url string, eventTypes []string) (response viberinterface.SetWebhookResponse, err error) {
	var responseBody []byte
	if _, responseBody, err = botApi.send(
		&viberinterface.SetWebhookMessage{
			ViberAuth: viberinterface.ViberAuth{Token: botApi.token},
			Url: url,
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

func (botApi ViberBotApi) SendMessage(m viberinterface.MessageToReceiver) (requestBody []byte, response viberinterface.SendMessageResponse, err error) {
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

func (botApi ViberBotApi) send(m viberinterface.MessageToViberEndpoint) (requestBody []byte, responseBody[]byte, err error) {
	if requestBody, err = ffjson.MarshalFast(m); err != nil {
		ffjson.Pool(requestBody)
		return
	}
	var resp *http.Response
	endpointUrl := VIBER_API_BASE_URL + m.Endpoint()
	resp, err = botApi.httpClient.Post(endpointUrl, "applicaiton/json", bytes.NewReader(requestBody));
	ffjson.Pool(requestBody)
	if err != nil {
		return
	}
	if responseBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}