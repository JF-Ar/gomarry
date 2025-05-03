package usersRepository

import (
	"github.com/JF-Ar/gomarry/schemas/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *user.User) error {
	return r.db.Create(u).Error
}
