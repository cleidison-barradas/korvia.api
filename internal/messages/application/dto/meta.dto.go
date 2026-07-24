package dto

type WebhookPayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string    `json:"messaging_product"`
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts"`
	Messages         []Message `json:"messages"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	Profile     Profile `json:"profile"`
	WaID        string  `json:"wa_id"`
	UserID      string  `json:"user_id"`
	CountryCode string  `json:"country_code"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From          string      `json:"from"`
	FromUserID    string      `json:"from_user_id"`
	ID            string      `json:"id"`
	Timestamp     string      `json:"timestamp"`
	Text          *MessageText `json:"text,omitempty"`
	FromLogicalID string      `json:"from_logical_id"`
	Type          string      `json:"type"`
}

type MessageText struct {
	Body string `json:"body"`
}