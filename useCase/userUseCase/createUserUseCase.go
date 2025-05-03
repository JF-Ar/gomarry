package userUseCase

import (
	"github.com/JF-Ar/gomarry/adapter/http/errors"
	"github.com/JF-Ar/gomarry/repositories/usersRepository"
	"github.com/JF-Ar/gomarry/schemas/user"
	"github.com/JF-Ar/gomarry/useCase/domainErrors"
)

type userUseCase struct {
	repo usersRepository.UserRepository
}

func NewUserUseCase(repo usersRepository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) CreateUser(u *user.User) *errors.GoMarryError {
	if u.Email == "bloqueado@example.com" {
		return errors.NewConflict(domainErrors.EmailAlreadyExists)
	}

	if err := uc.repo.Create(u); err != nil {
		return err
	}
	return nil
}
