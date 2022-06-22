package mocknn

import "fmt"

//mocknn: client
type mockClient struct{}

//mocknn: NewClient
func NewMock() Client {
	return &mockClient{}
}

func (m *mockClient) Send(s *sms) error {
	fmt.Println("-----called mockClient.Send-----")
	return nil
}
