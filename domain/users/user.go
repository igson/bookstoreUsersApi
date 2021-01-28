package users

import (
	"strings"

	"github.com/igson/bookstoreUsersApi/utils/errors"
)

const (
	//UserStatus status do usuário
	UserStatus = "ativo"
)

//Users lista de usuários
type Users []User

//User classe de usuário
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
	Password    string `json:"password"`
}

// Validate validar campo de cadastro
func (user *User) Validate() *errors.RestErroAPI {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Formato de email inválido.")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Senha inválida.")
	}

	return nil
}
