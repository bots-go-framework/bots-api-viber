package viberbotapi

import "net/http"

// BotAPI allows you to interact with the Telegram Bot API.
type ViberBotApi struct {
	httpClient http.Client
	token      string
}

func NewViberBotApiWithHttpClient(token string, httpClient http.Client) ViberBotApi {
	return ViberBotApi{token: token, httpClient: httpClient}
}

func (botApi ViberBotApi) SetWebhook(url string, eventTypes []string) error {
	m := SetWebhookMessage{
		Url: url,
		EventTypes: eventTypes,
	}
	m.Token = botApi.token
	return botApi.SendMessage(m)
}

func (bot ViberBotApi) SendMessage(m MessageToViber) error {
	return nil
}