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
	var body []byte
	if body, err = botApi.send(
		&viberinterface.SetWebhookMessage{
			ViberAuth: viberinterface.ViberAuth{Token: botApi.token},
			Url: url,
			EventTypes: eventTypes,
		},
	); err != nil {
		return
	}

	if err = ffjson.UnmarshalFast(body, &response); err != nil {
		err = errors.Wrap(err, "Failed to unmarshal response body to JSON")
		return
	}
	return
}

func (botApi ViberBotApi) SendMessage(m viberinterface.MessageToReceiver) (response viberinterface.SendMessageResponse, err error) {
	m.SetType(m.GetType())
	var body []byte
	if body, err = botApi.send(m); err != nil {
		return
	}

	if err = ffjson.UnmarshalFast(body, &response); err != nil {
		err = errors.Wrap(err, "Failed to unmarshal response body to JSON")
		return
	}
	return response, nil
}

func (bot ViberBotApi) send(m viberinterface.MessageToViberEndpoint) ([]byte, error) {
	endpointUrl := VIBER_API_BASE_URL + m.Endpoint()
	body, err := ffjson.MarshalFast(m)
	if err != nil {
		return nil, err
	}
	resp, err := bot.httpClient.Post(endpointUrl, "applicaiton/json", bytes.NewReader(body))
	//if resp != nil {
	//	//defer func() {
	//	//	resp.Body.Close()
	//	//}()
	//	//if resp.Request != nil {
	//	//	defer func() {
	//	//		resp.Request.Body.Close()
	//	//	}()
	//	//}
	//}
	if err != nil {
		return nil, err
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return body, err
}