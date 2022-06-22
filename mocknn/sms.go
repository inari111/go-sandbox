package mocknn

// SMS Client
type Client interface {
	Send(s *sms) error
}

type sms struct {
	PhoneNumber string
	Message     string
}

func NewClient() Client {
	return &client{}
}

type client struct{}

func (c *client) Send(s *sms) error {
	panic("implement me")
}
