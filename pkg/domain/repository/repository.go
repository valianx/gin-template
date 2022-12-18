package repository

import (
	"fmt"
	"github.com/valianx/gin-template/pkg/domain/model"
	"gorm.io/gorm"
)

func (r *UserRepositoryImpl) FindById(id int) (*model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
