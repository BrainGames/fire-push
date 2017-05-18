package fire_push

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type firebaseClient struct {
	serverKey  string
	timeToLive int
	color      string
	sound      string
}

func NewFirebaseClient(key string, timeToLive int, color, sound string) *firebaseClient {
	if sound == "" {
		sound = "default"
	}
	return &firebaseClient{serverKey: key, timeToLive: timeToLive, color: color, sound: sound}
}

func (fc *firebaseClient) SendData(recipientKey, title, body, image string, badge int) (*http.Response, error) {
	payload := &payload{}
	payload.To = recipientKey
	pd := &PayloadData{Title: title, Body: body, Image:image, Badge: badge, Color: fc.color, Sound: fc.sound}
	payload.Data = pd
	p, _ := json.Marshal(payload)
	c := &http.Client{}
	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(p))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+fc.serverKey)
	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil
}

type payload struct {
	To           string               `json:"to"`
	Data         *PayloadData         `json:"data"`
	Notification *PayloadNotification `json:"notification"`
}
type PayloadData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
	Badge int    `json:"badge"`
	Color string `json:"color"`
	Sound string `json:"sound"`
}

type PayloadNotification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Badge int    `json:"badge"`
	Color string `json:"color"`
	Sound string `json:"sound"`
}

func (fc *firebaseClient) SendNotification(recipientKey, title, body string, badge int) (*http.Response, error) {
	payload := &payload{}
	payload.To = recipientKey
	pn := &PayloadNotification{Title: title, Body: body, Badge: badge, Color: fc.color, Sound: fc.sound}
	payload.Notification = pn
	p, _ := json.Marshal(payload)
	c := &http.Client{}
	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(p))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+fc.serverKey)
	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil
}
