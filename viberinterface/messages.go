package viberinterface

const(
	endpointSetWebhook = "set_webhook"
	endpointSendMessage = "send_message"
)

type MessageToViberEndpoint interface {
	isMessageToViberEndpoint()
	Endpoint() string
}

type MessageToReceiver interface {
	MessageToViberEndpoint
	isMessageToReceiver()
	GetType() string
	SetType(mType string)
	SetToken(token string)
}

type isMessageToViber struct {
}

func (_ isMessageToViber) isMessageToViberEndpoint() {
}

func (_ BaseMessageToReceiver) isMessageToReceiver() {
}


type ViberAuth struct {
	Token string `json:"auth_token"`
}

var _ MessageToViberEndpoint = (*SetWebhookMessage)(nil)

type SetWebhookMessage struct {
	isMessageToViber
	ViberAuth
	Url        string `json:"url"`
	EventTypes []string `json:"event_types,omitempty"`
}

func (_ SetWebhookMessage) Endpoint() string {
	return endpointSetWebhook
}

type SetWebhookResponse struct {
	Status        int `json:"status"`
	StatusMessage string `json:"status_message"`
	EventTypes    []string `json:"event_types,omitempty"`
}

type WebhookCallback struct {
	Event        string `json:"event"`         // Callback type – which event triggered the callback: webhook
	Timestamp    int `json:"timestamp"`        // Epoch time of the event that triggered the callback
	MessageToken int64 `json:"message_token"`  // Unique ID of the message
}

type PaSender struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type BaseMessageToReceiver struct {
	ViberAuth
	Receiver     string `json:"receiver"`
	Sender       *PaSender `json:"sender,omitempty"`
	MessageType  string `json:"type"`
	Keyboard     *Keyboard `json:"keyboard,omitempty"`
	TrackingData string `json:"tracking_data"`
}

func (_ BaseMessageToReceiver) Endpoint() string {
	return endpointSendMessage
}

func (m *BaseMessageToReceiver) SetType(mType string) {
	m.MessageType = mType
}

func (m *BaseMessageToReceiver) SetToken(token string) {
	m.ViberAuth = ViberAuth{Token: token}
}

var _ MessageToReceiver = (*TextMessage)(nil)

type TextMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	Text string `json:"text"`
}

func (_ TextMessage) GetType() string {
	return "text"
}

func NewTextMessage(receiver, trackingData, text string, keyboard *Keyboard) *TextMessage {
	return &TextMessage{
		Text: text,
		BaseMessageToReceiver: BaseMessageToReceiver{
			Receiver: receiver,
			TrackingData: trackingData,
			Keyboard: keyboard,
		},
	}
}

var _ MessageToReceiver = (*PictureMessage)(nil)

type PictureMessage struct {
	TextMessage
	Media     string `json:"media"`
	Thumbnail string `json:"thumbnail"`
}

func (_ PictureMessage) GetType() string {
	return "picture"
}

var _ MessageToReceiver = (*VideoMessage)(nil)

type VideoMessage struct {
	isMessageToViber
	PictureMessage
	Size     int `json:"size"`
	Duration int `json:"duration"`
}

func (_ VideoMessage) GetType() string {
	return "video"
}

var _ MessageToReceiver = (*FileMessage)(nil)

type FileMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	Media    string `json:"media"`
	FileName string `json:"file_name"`
	Size     int `json:"size"`
}

func (_ FileMessage) GetType() string {
	return "file"
}

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

var _ MessageToReceiver = (*ContactMessage)(nil)

type ContactMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	Contact Contact `json:"contact"`
}

func (_ ContactMessage) GetType() string {
	return "contact"
}

type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

var _ MessageToReceiver = (*LocationMessage)(nil)

type LocationMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	Location Location `json:"location"`
}

func (_ LocationMessage) GetType() string {
	return "location"
}

var _ MessageToReceiver = (*UrlMessage)(nil)

type UrlMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	Media string `json:"media"`
}

func (_ UrlMessage) GetType() string {
	return "url"
}

var _ MessageToReceiver = (*StickerMessage)(nil)

type StickerMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	StickerID string `json:"sticker_id"`
}

func (_ StickerMessage) GetType() string {
	return "sticker"
}

var _ MessageToReceiver = (*KeyboardMessage)(nil)

type KeyboardMessage struct {
	isMessageToViber
	BaseMessageToReceiver
	//Keyboard Keyboard `json:"keyboard"`
}

func NewKeyboardMessage(receiver, trackingData string, keyboard *Keyboard) *KeyboardMessage {
	if keyboard == nil {
		panic("keyboard == nil")
	}
	return &KeyboardMessage{
		BaseMessageToReceiver: BaseMessageToReceiver{
			Receiver: receiver,
			TrackingData: trackingData,
			Keyboard: keyboard,
		},
	}
}

func (_ KeyboardMessage) GetType() string {
	return "keyboard"
}


// https://developers.viber.com/customer/en/portal/articles/2632255-send-message?b_id=15145#send-message-response
type SendMessageResponse struct {
	Status        int `json:"status"`           // 0 for success, otherwise failure
	StatusMessage string `json:"status_message"`
	MessageToken  int64 `json:"message_token"` // Unique ID of the message
}