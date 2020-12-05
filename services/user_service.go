package services

import (
	"github.com/igson/bookstoreUsersApi/domain/users"
)

//CreateUser camada de regra de négocio pra cadastro de usuários
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
