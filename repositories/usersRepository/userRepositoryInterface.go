package usersRepository

import (
	"github.com/JF-Ar/gomarry/adapter/http/errors"
	"github.com/JF-Ar/gomarry/schemas/user"
)

type UserRepository interface {
	Create(u *user.User) *errors.GoMarryError
}
