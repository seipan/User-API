package repository

import (
	"User-API/domain/entity"
	"User-API/domain/repository"
	db_error "User-API/error/db"
	"context"
	"database/sql"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	statement := "INSERT INTO users VALUES($1,$2,$3)"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, db_error.StatementError
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Id, user.Name, user.Mail)

	if err != nil {
		log.Println(err)
		return nil, db_error.ExecError
	}

	resuser := &entity.User{}
	resuser.Id = user.Id
	resuser.Name = user.Name
	resuser.Mail = user.Mail

	return resuser, nil
}

func (ur userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	statement := "UPDATE users SET name = $2, mail = $3 WHERE id = $1"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, db_error.StatementError
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Id, user.Name, user.Mail)

	if err != nil {
		log.Println(err)
		return nil, db_error.ExecError
	}

	resuser := &entity.User{}
	resuser.Id = user.Id
	resuser.Name = user.Name
	resuser.Mail = user.Mail

	return resuser, nil
}

func (ur userRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	statement := "SELECT * FROM users WHERE id = $1"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, db_error.StatementError
	}
	defer stmt.Close()
	resuser := &entity.User{}

	err = stmt.QueryRowContext(ctx, id).Scan(&resuser.Id, &resuser.Name, &resuser.Mail)

	if err != nil {
		log.Println(err)
		return nil, db_error.QueryError
	}

	return resuser, nil
}

func (ur userRepository) DeleteUser(ctx context.Context, id string) error {
	statement := "DELETE FROM users WHERE id = $1"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return db_error.StatementError
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		log.Println(err)
		return db_error.ExecError
	}

	return nil
}
