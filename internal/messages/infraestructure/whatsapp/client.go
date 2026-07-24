package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)


type Client struct {
	PhoneNumberID string
	AccessToken string
	BaseURL string
}

func NewClient(phoneNumberID string, accessToken string, baseURL string) *Client {
	return &Client{
		PhoneNumberID: "1113112031894761",
		AccessToken: "EAAWEKAUghMcBSPPO7jnJZCITMJ9fzJtwr2HQONbbIm8K89Rc8mRLeyunwAveEesClCXxrTHU9gppdwbP08mFXnm01LwZCam9eOH8vGYgeAQQn3R8EhGJDcP1ZCrqnsevVZCGJhaK6jcwOC3o7DKxwhQ2IqWivo96tDwCILzXa8MnyjBOaelyL31sZAkNZC5vPCZAmOLV1grpreHiX5ERPGZCwajhNhr8ApeZAHYW3bzBEbxelVLS3a4PaGaj0wyHK2rRrDCZCkw1Y5t2WLJ33dWqxyeBQm",
		BaseURL: "https://graph.facebook.com/v25.0",
	}
}

type textMessagePayload struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

func (c *Client) SendText(to string, text string) error {
	payload := textMessagePayload{
		MessagingProduct: "whatsapp",
		To:               to,
		Type:             "text",
	}
	payload.Text.Body = text

	return c.send(payload)
}

func (c *Client) SendButtons(to string, text string, options []string) error {
	full := text

	for i, opt := range options {
		full += fmt.Sprintf("\n%d. %s", i+1, opt)
	}
	
	return c.SendText(to, full)
}

func (c *Client) send(payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s/messages", c.BaseURL, c.PhoneNumberID)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))

	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode, resp.Body)

	if resp.StatusCode >= 300 {
	var errBody bytes.Buffer
	errBody.ReadFrom(resp.Body)
	fmt.Println("WHATSAPP ERROR BODY:", errBody.String()) // <- isso aqui mostra o JSON real
	return fmt.Errorf("whatsapp api error [%d]: %s", resp.StatusCode, errBody.String())
	}

	return nil
}