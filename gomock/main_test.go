package gomock

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_userRepo_Create(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name        string
		prepareMock func(m *MockUserRepository, u *User)
		args        args
		want        *User
		wantErr     bool
	}{
		{
			name: "正常系",
			prepareMock: func(m *MockUserRepository, u *User) {
				m.EXPECT().Create(u).Return(&User{
					ID:   u.ID,
					Name: u.Name,
				}, nil)
			},
			args: args{
				u: &User{
					ID:   "1234abcd",
					Name: "hoge",
				},
			},
			want: &User{
				ID:   "1234abcd",
				Name: "hoge",
			},
			wantErr: false,
		},
		{
			name: "エラーを返す場合",
			prepareMock: func(m *MockUserRepository, u *User) {
				m.EXPECT().Create(u).Return(nil, errors.New("new error"))
			},
			args: args{
				u: &User{
					ID:   "1234abcd",
					Name: "hoge",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// モック生成
			mockRepo := NewMockUserRepository(ctrl)

			// モックがどういう値を返すか決める
			tt.prepareMock(mockRepo, tt.args.u)

			// MockUserRepository.Create() の呼び出し
			got, err := mockRepo.Create(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// モックが返した値と期待値の比較
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("diff (-got +want)\n%s", diff)
			}
		})
	}
}
