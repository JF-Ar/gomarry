package userHandler

import (
	"github.com/JF-Ar/gomarry/adapter/http/errors"
	"github.com/JF-Ar/gomarry/config"
	"github.com/JF-Ar/gomarry/helper"
	"github.com/JF-Ar/gomarry/infra/requests/userRequestHandler"
	"github.com/JF-Ar/gomarry/schemas/user"
	"github.com/JF-Ar/gomarry/useCase/userUseCase"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	logger *config.Logger
)

type Handler struct {
	createUC userUseCase.UserUseCase
}

func NewHandler(uc userUseCase.UserUseCase) *Handler {
	return &Handler{createUC: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var userRequest userRequestHandler.CreateUserRequest
	userFields, err := userRequest.NewCreateUser(c)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	if err = userFields.Validate(); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	domainUser, err := userFieldsToDomain(userFields)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	if err = h.createUC.CreateUser(domainUser); err != nil {
		internal := errors.NewInternal("Internal Server Error")
		c.JSON(internal.StatusCode, internal)
		return
	}

	c.JSON(http.StatusCreated, "User created successfully")

}

func userFieldsToDomain(userFields *userRequestHandler.CreateUserRequest) (*user.User, *errors.GoMarryError) {
	logger = config.GetLogger("UserHandler")

	password, err := helper.HashPassword(userFields.Password)
	if err != nil {
		logger.ErrorF("Error hashing password: ", err)
		return nil, errors.NewInternal("Internal Server Error")
	}

	return &user.User{
		Username: userFields.Name,
		Password: password,
		Email:    userFields.Email,
		Phone:    userFields.Phone,
		Role:     userFields.Role,
	}, nil
}
