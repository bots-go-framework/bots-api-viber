package viberbotapi

import "net/http"

// BotAPI allows you to interact with the Telegram Bot API.
type ViberBotApi struct {
	httpClient http.Client
	Token  string
}

func NewViberBotApiWithClient(token string, httpClient http.Client) ViberBotApi {
	return ViberBotApi{Token: token, httpClient: httpClient}
}

func (botApi ViberBotApi) SetWebhook(url string, eventTypes []string) error {
	m := SetWebhookMessage{
		Url: url,
		EventTypes: eventTypes,
	}
	m.Token = botApi.Token
	return botApi.SendMessage(m)
}

func (bot ViberBotApi) SendMessage(m MessageToViber) error {
	return nil
}