package fp

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

func (fc firebaseClient) SendData(recipient, title, body, image string, badge int) error {
	payload := &Payload{
		To: recipient,
		Data: &PayloadData{
			Title: title,
			Body:  body,
			Image: image,
			Badge: badge,
			Color: fc.color,
			Sound: fc.sound,
		},
	}

	return fc.processRequest(payload)
}

func (fc firebaseClient) SendNotification(recipient, title, body string, badge int) error {
	payload := &Payload{
		To: recipient,
		Notification: &PayloadNotification{
			Title: title,
			Body:  body,
			Badge: badge,
			Color: fc.color,
			Sound: fc.sound,
		},
	}

	return fc.processRequest(payload)
}

func (fc firebaseClient) processRequest(p *Payload) error {
	body, _ := json.Marshal(p)

	c := http.DefaultClient

	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+fc.serverKey)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
