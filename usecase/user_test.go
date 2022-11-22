package usecase

import (
	"User-API/domain/entity"
	"User-API/error/usecase"
	"context"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		reqname string
		reqmail string
		wantErr error
	}{
		{
			name:    "正常に動作した場合",
			reqname: "hoge",
			reqmail: "hoge@hoge.com",
			wantErr: nil,
		},
		{
			name:    "nameが空なら name empty error",
			reqname: "",
			reqmail: "hoge@hoge.com",
			wantErr: usecase.NameEmptyError,
		},
		{
			name:    "mailが空なら  empty error",
			reqname: "hoge",
			reqmail: "",
			wantErr: usecase.MailEmptyError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userUsecase := NewUserUsecase(&UserRepositoryMock{})

			user := &entity.User{}
			user.Name = tt.reqname
			user.Mail = tt.reqmail
			ctx := context.Background()

			resuser, err := userUsecase.CreateUser(ctx, user)

			if err != tt.wantErr {
				t.Errorf("TestHandler_CreateTask code Error : want %s but got %s", tt.wantErr, err)
			}

			if resuser != nil && resuser.Mail != user.Mail {
				t.Errorf("TestHandler_CreateTask code Error : want %s but got %s", user.Mail, resuser.Mail)
			}

			if resuser != nil && resuser.Name != user.Name {
				t.Errorf("TestHandler_CreateTask code Error : want %s but got %s", user.Name, resuser.Name)
			}
		})
	}
}

type UserRepositoryMock struct{}

var (
	testCreateUser *entity.User = &entity.User{
		Id:   1,
		Name: "hoge",
		Mail: "hoge@hoge.com",
	}
)

func (h *UserRepositoryMock) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserRepositoryMock) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserRepositoryMock) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserRepositoryMock) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
