package fire_push

import (
	"net/http"
	"bytes"
	"encoding/json"
)

type androidClient struct {
	serverKey string
	timeToLive int64
}

func NewAndroidClient (key string) *androidClient {
	return &androidClient{serverKey:key}
}
func (ac *androidClient) SetTimeToLive(value int64) {
	ac.timeToLive = value
}
func (ac *androidClient) Send (recipientKey, title, body, image string, badge int) (*http.Response, error){
	payload := &androidPayload{}
	payload.To = recipientKey
	payload.Data.Title = title
	payload.Data.Body = body
	payload.Data.Image = image
	payload.Data.Badge = badge
	p, _ := json.Marshal(payload)
	c := &http.Client{}
	req, _ := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewReader(p))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+ac.serverKey)
	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil

}

type androidPayload struct {
	To string `json:"to"`
	Data struct{
		Title string `json:"title"`
		Body string `json:"body"`
		Image string `json:"image"`
		Badge int `json:"badge"`
	} `json:"data"`
}