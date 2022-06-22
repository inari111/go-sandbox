package mocknn

import "testing"

func Test_client_Send(t *testing.T) {
	type args struct {
		s *sms
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				s: &sms{
					PhoneNumber: "09012345678",
					Message:     "test",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{}
			if err := c.Send(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
