package repository

import (
	"gorm.io/gorm"
)

/*
type UserRepository interface {
	FindById(id int) (*model.User, error)
	Save(user *model.User) error
}
*/

type UserRepositoryImpl struct {
	DB *gorm.DB
}
