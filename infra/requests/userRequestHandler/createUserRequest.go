package userRequestHandler

import (
	"github.com/JF-Ar/gomarry/adapter/http/errors"
	"github.com/JF-Ar/gomarry/helper/constants"
	"github.com/gin-gonic/gin"
	"regexp"
	"slices"
	"strings"
)

var sliceRoles = []string{"groom", "bride", "support"}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,phone"`
	Password string `json:"password" binding:"required,password"`
	Role     string `json:"role" binding:"required,role"`
}

func (c *CreateUserRequest) NewCreateUser(context *gin.Context) (*CreateUserRequest, *errors.GoMarryError) {
	fieldsUser := &CreateUserRequest{}

	if err := context.ShouldBind(fieldsUser); err != nil {
		return nil, errors.NewBadRequest(err.Error())
	}

	return fieldsUser, nil
}

func (c *CreateUserRequest) Validate() *errors.GoMarryError {
	if err := c.validatePassword(); err != nil {
		return err
	}

	if err := c.validateEmail(); err != nil {
		return err
	}

	if err := c.validateRole(); err != nil {
		return err
	}

	return nil
}

func (c *CreateUserRequest) validatePassword() *errors.GoMarryError {
	if len(c.Password) < 8 {
		return errors.NewBadRequest(constants.PasswordTooShortErrorMessage)
	}

	return nil
}

func (c *CreateUserRequest) validateEmail() *errors.GoMarryError {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !re.MatchString(c.Email) {
		return errors.NewBadRequest(constants.InvalidFormatEmail)
	}

	return nil
}

func (c *CreateUserRequest) validateRole() *errors.GoMarryError {
	re := strings.ToLower(c.Role)

	if !slices.Contains(sliceRoles, re) {
		return errors.NewBadRequest(constants.InvalidRole)
	}

	return nil
}
