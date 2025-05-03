package userUseCase

import (
	"github.com/JF-Ar/gomarry/adapter/http/errors"
	"github.com/JF-Ar/gomarry/schemas/user"
)

type UserUseCase interface {
	CreateUser(*user.User) *errors.GoMarryError
}
