package viberbotapi

type MessageToViber interface {
}

type ViberAuth struct {
	Token string `json:"auth_token"`
}

type SetWebhookMessage struct {
	ViberAuth
	Url        string `json:"url"`
	EventTypes []string `json:"event_types,omitempty"`
}

type SetWebhookResponse struct {
	Status        int `json:"status"`
	StatusMessage string `json:"status_message"`
	EventTypes    []string `json:"event_types,omitempty"`
}

type WebhookCallback struct {
	Event        string `json:"event"`  // Callback type – which event triggered the callback: webhook
	Timestamp    int `json:"timestamp"` // Epoch time of the event that triggered the callback
	MessageToken string `json:"event"`  // Unique ID of the message
}

type Sender struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ViberBaseMessage struct {
	ViberAuth
	Receiver string `json:"receiver"`
	Type     string `json:"type"`
}

type TextMessage struct {
	ViberBaseMessage
	Text         string `json:"name"`
	TrackingData string `json:"tracking_data"`
}

type PictureMessage struct {
	TextMessage
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
}

type VideoMessage struct {
	PictureMessage
	Size     int `json:"size"`
	Duration int `json:"duration"`
}

type FileMessage struct {
	ViberBaseMessage
	Media    string `json:"media"`
	FileName string `json:"file_name"`
	Size     int `json:"size"`
}

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type ContactMessage struct {
	ViberBaseMessage
	Contact Contact `json:"contact"`
}

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type LocationMessage struct {
	ViberBaseMessage
	Location Location `json:"location"`
}

type UrlMessage struct {
	ViberBaseMessage
	Media    string `json:"media"`
}

type StickerMessage struct {
	ViberBaseMessage
	StickerID    string `json:"sticker_id"`
}

// https://developers.viber.com/customer/en/portal/articles/2632255-send-message?b_id=15145#send-message-response
type SendMessageResponse struct {
	MessageToken string `json:"message_token"` // Unique ID of the message
	Status        int `json:"status"` // 0 for success, otherwise failure
	StatusMessage string `json:"status_message"`
}