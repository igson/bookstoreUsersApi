package services

import (
	"fmt"
	"strings"

	"github.com/igson/bookstoreUsersApi/domain/users"
	"github.com/igson/bookstoreUsersApi/utils/errors"
)

const (
	uniqueEmailConstraint = "email_UNIQUE"
)

//GetUser buscar usuario pelo ID
func GetUser(userID int64) (*users.User, *errors.RestErroAPI) {

	user := &users.User{ID: userID}

	if erro := user.GetUser(); erro != nil {
		return nil, erro
	}

	return user, nil

}

//CreateUser camada de regra de négocio pra cadastro de usuários
func CreateUser(user users.User) (*users.User, *errors.RestErroAPI) {

	if erro := user.Validate(); erro != nil {
		return nil, erro
	}

	if erro := user.Save(); erro != nil {

		if strings.Contains(erro.Message, uniqueEmailConstraint) {
			return nil, errors.NewInternalServerError(fmt.Sprintf("Email %s já cadastrado:", user.Email))
		}

		return nil, erro
	}

	return &user, nil
}
