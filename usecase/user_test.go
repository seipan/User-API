package usecase

import (
	"User-API/domain/entity"
	"context"
	"net/http"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	tests := []struct {
		name      string
		reqBody   string
		reqMethod string
		reqname   string
		reqmail   string
		wantCode  int
		wantBody  string
	}{
		{
			name:      "正常に動作した場合",
			reqBody:   `{"name":"hoge","mail":"hoge@hoge.com"}`,
			wantCode:  http.StatusOK,
			reqname:   "hoge",
			reqmail:   "hoge@hoge.com",
			reqMethod: http.MethodPost,
			wantBody:  `{"name":"hoge","id":1,"mail":"hoge@hoge.com"}`,
		},
		{
			name:      "request bodyが空だった場合、400エラーになる",
			reqBody:   ``,
			wantCode:  200,
			reqname:   "",
			reqmail:   "",
			reqMethod: http.MethodPost,
			wantBody:  `{"Status":400,"Result":"name empty"}`,
		},
		{
			name:      "POSTリクエスト以外は 404 ",
			reqBody:   `{"name":"hoge","mail":"hoge@hoge.com"}`,
			wantCode:  405,
			reqname:   "hoge",
			reqmail:   "hoge@hoge.com",
			reqMethod: http.MethodGet,
			wantBody:  `{"Status":404,"Result":"method not allowed"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.wantCode != w.Code {
				t.Errorf("TestHandler_CreateTask code Error : want %d but got %d", tt.wantCode, w.Code)
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
