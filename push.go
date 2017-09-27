package fp

type Sender interface {
	SendNotification(recipient, title, body string, badge int) error
	SendData(recipient, title, body, image string, badge int) error
}

type Payload struct {
	To           string               `json:"to"`
	Data         *PayloadData         `json:"data,omitempty"`
	Notification *PayloadNotification `json:"notification,omitempty"`
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

var sender Sender

func Init(s Sender) {
	sender = s
}

func SendNotification(recipient, title, body string, badge int) error {
	return sender.SendNotification(recipient, title, body, badge)
}

func SendData(recipient, title, body, image string, badge int) error {
	return sender.SendData(recipient, title, body, image, badge)
}
