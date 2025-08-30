package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"ms-user/internal/model"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user model.User) (model.User, error) {
	var created model.User

	query := `
             INSERT INTO users (username, email, password) 
             VALUES ($1, $2, $3)
             RETURNING id, username, email, password
    `

	err := r.db.Get(&created, query, user.Username, user.Email, user.Password)

	if err != nil {
		return model.User{}, err
	}

	return created, nil
}

func (r *UserRepositoryImpl) Delete(id int64) error {
	res, err := r.db.Exec("delete from users where id=$1", id)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepositoryImpl) GetAll() ([]model.User, error) {
	var users []model.User

	err := r.db.Select(&users, "select id, username, email, password from users")

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetByID(id int64) (model.User, error) {
	var user model.User

	err := r.db.Get(&user, "select id, username, email, password from users where id=$1", id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
