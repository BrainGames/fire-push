package fire_push

type iosClient struct {
	serverKey string
}

func NewIOSClient (key string) *iosClient {
	return &iosClient{serverKey:key}
}
