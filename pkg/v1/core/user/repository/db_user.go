package repository

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Register() error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&datamodel.User{})
	return &UserRepository{db}
}

func (r *UserRepository) Register() error {
	return nil
}
