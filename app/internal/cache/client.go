package cache

type Client struct {
	client Cacher
}

func NewClient(client Cacher) *Client {
	return &Client{client: client}
}
