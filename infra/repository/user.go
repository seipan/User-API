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
	statement := "INSERT INTO users (name,mail) VALUES($1,$2)"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.Name, user.Mail)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	resId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resuser := &entity.User{}
	resuser.Id = resId
	resuser.Name = user.Name
	resuser.Mail = user.Mail

	return resuser, nil
}

func (ur userRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	statement := "UPDATE users SET name = $2, mail = $3 WHERE id = $1"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.Id, user.Name, user.Mail)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	resId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resuser := &entity.User{}
	resuser.Id = resId
	resuser.Name = user.Name
	resuser.Mail = user.Mail

	return resuser, nil
}

func (ur userRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	statement := "SELECT * FROM users WHERE id = $1"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(db_error.StatementError)
		return nil, err
	}
	defer stmt.Close()
	resuser := &entity.User{}

	err = stmt.QueryRow(id).Scan(&resuser.Id, &resuser.Name, &resuser.Mail)

	if err != nil {
		log.Println(db_error.QueryError)
		return nil, db_error.QueryError
	}

	return resuser, nil
}

func (ur userRepository) DeleteUser(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}
