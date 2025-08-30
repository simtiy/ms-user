package repository

import (
	"ms-user/internal/model"
)

type IUserRepository interface {
	GetByID(id int64) (model.User, error)
	GetAll() ([]model.User, error)
	Create(user model.User) (model.User, error)
	Delete(id int64) error
}
