package repository

import (
	"User-API/domain/entity"
	"context"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_UserCreate(t *testing.T) {
	log.Println("ouou")

	t.Run(
		"Createが成功する",
		func(t *testing.T) {
			// Arrange
			testuser := &entity.User{
				Name: "testName",
				Mail: "test@test.com",
			}
			db_test, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db_test.Close()
			mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO users")).
				ExpectExec().
				WithArgs(testuser.Name, testuser.Mail).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := &userRepository{db: db_test}
			ctx := context.Background()
			_, err = r.CreateUser(ctx, testuser)

			if err != nil {
				t.Error(err.Error())
			}
		},
	)

}
