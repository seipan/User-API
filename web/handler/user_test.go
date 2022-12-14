package handler

import (
	"User-API/domain/entity"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
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
			mux := http.NewServeMux()

			userHandler := NewUserhandler(&UserUsecaseMock{})
			mux.Handle("/user/create", http.HandlerFunc(userHandler.CreateUser))

			//body := bytes.NewBufferString(tt.reqBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.reqMethod, "/user/create", nil)

			values := url.Values{}

			values.Set("name", tt.reqname)
			values.Add("mail", tt.reqmail)
			req.URL.RawQuery = values.Encode()

			mux.ServeHTTP(w, req)

			if tt.wantCode != w.Code {
				t.Errorf("TestHandler_CreateTask code Error : want %d but got %d", tt.wantCode, w.Code)
			}

			wantbody := tt.wantBody + "\n"

			if tt.wantBody != w.Body.String() && wantbody != w.Body.String() {
				t.Errorf("TestHandler_CreateTask body Error : want %s but got %s", tt.wantBody, w.Body.String())
			}

		})
	}
}

//usecase の mock を書く

var (
	testCreateUser *entity.User = &entity.User{
		Id:   1,
		Name: "hoge",
		Mail: "hoge@hoge.com",
	}
)

type UserUsecaseMock struct{}

func (h *UserUsecaseMock) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
