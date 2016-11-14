package viberinterface


type CallbackBase struct {
	Event   string `json:"event"`
	Timestamp int `json:"timestamp"`
}

type CallbackOnMessage struct {
	MessageToken int64 `json:"message_token"`
	Sender CallbackSender `json:"sender"`
	Message CallbackMessage `json:"message"`
}

type CallbackSender struct {
	ID string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type CallbackMessage struct {
	Type         string `json:"type"`
	TrackingData string `json:"tracking_data"`
	Text string `json:"text"`
	Media string `json:"media"`
	Location *Location `json:"location"`
	Contact *Contact `json:"contact"`
}

type CallbackUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Country  string `json:"country"`
	Language string `json:"language"`
}

type CallbackOnSubscribed struct {
	CallbackBase
	User CallbackUser `json:"user"`
}

type CallbackOnUserIdBase struct {
	CallbackBase
	UserID string `json:"user_id"`
}

type CallbackOnUnsubscribed CallbackOnUserIdBase
type CallbackOnDelivered CallbackOnUserIdBase
type CallbackOnSeen CallbackOnUserIdBase

type CallbackOnFailed struct {
	CallbackOnUserIdBase
	Description string `json:"desc"`
}

/*  https://developers.viber.com/customer/en/portal/articles/2541267-callbacks?b_id=15145#conversation-started

Conversation opened event fires when a user opens a conversation with the PA for the first time (no conversation history with the PA), or when the conversation is opened using a deep link (See Deep link section for more information).

This event is not considered a subscribe event and doesn’t allow the PA to send messages to the user; however, it will allow sending one “welcome message” to the user.

Once a conversation_opened callback is received, the service will be able to respond with aJson containing same parameters as a send_message request. "auth_token" and "receiver" are not mandatory in this case.

Example response: {"sender": { "name": "yarden from the pa", "avatar": "http://avatar_url",}, tracking_data": "tracking data", "type": "picture", "text": "Photo description", "media":"http://www.images.com/img.jpg", "thumbnail": "http://www.images.com/thumb.jpg" }

 */
type CallbackOnConversationStarted struct {
	CallbackBase
	MessageToken int64 `json:"message_token"`
	Type         string `json:"type"` // "open". Additional types may be added in the future.
	Context      string `json:"context"`
	User         CallbackUser `json:"user"`
}

type Callback struct {
	CallbackBase
}