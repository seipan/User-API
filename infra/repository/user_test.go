package repository

import (
	"User-API/domain/entity"
	db_error "User-API/error/db"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserCreate(t *testing.T) {
	tests := []struct {
		name    string
		id      int64
		mail    string
		wantErr error
	}{
		{
			name:    "存在しないIDは作成できる",
			id:      1,
			mail:    "test@test.com",
			wantErr: nil,
		},
		{
			name:    "存在するIDは作成できない",
			id:      1,
			mail:    "test@test.com",
			wantErr: db_error.QueryError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, _, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()

			r := &userRepository{db: db}

			testuser := &entity.User{}
			testuser.Id = tt.id
			testuser.Name = tt.name
			testuser.Mail = tt.mail
			ctx := context.TODO()

			if _, err := r.CreateUser(ctx, testuser); !errors.Is(err, tt.wantErr) {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
