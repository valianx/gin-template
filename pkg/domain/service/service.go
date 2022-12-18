package service

import (
	"github.com/valianx/gin-template/pkg/domain/repository"
	"github.com/valianx/gin-template/pkg/infra/database"
)

func GetUserName(id int) string {
	db := database.DB
	userRepo := &repository.UserRepositoryImpl{DB: db}
	user, err := userRepo.FindById(id)

	if err != nil || user == nil {
		return ""
	}
	return user.Name
}
