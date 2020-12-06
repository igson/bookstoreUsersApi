package users

import (
	"fmt"
	"log"

	"github.com/igson/bookstoreUsersApi/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

//GetUser retorna usuário pelo id
func (user *User) GetUser() *errors.RestErroAPI {
	usuarioEncontrado := usersDB[user.ID]
	if usuarioEncontrado == nil {
		return errors.NewNotFoundErro(fmt.Sprintf("Usuário %d não encontado", user.ID))
	}
	user.ID = usuarioEncontrado.ID
	user.Email = usuarioEncontrado.Email
	user.FirstName = usuarioEncontrado.FirstName
	user.LastName = usuarioEncontrado.LastName
	user.DateCreated = usuarioEncontrado.DateCreated

	return nil
}

//Save cadastrar novo usuário
func (user *User) Save() *errors.RestErroAPI {

	u := usersDB[user.ID]

	log.Printf("Email  %s", user.Email)

	if u != nil {
		log.Printf("Email  %s", u.Email)
		if u.Email == user.Email {
			fmt.Println("Email já cadastrado")
			return errors.NewBadRequestError(fmt.Sprintf("Email %s já cadastrado", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d já existe", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
