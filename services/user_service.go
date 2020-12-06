package services

import (
	"github.com/igson/bookstoreUsersApi/domain/users"
	"github.com/igson/bookstoreUsersApi/utils/errors"
)

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
		return nil, erro
	}

	return &user, nil
}
