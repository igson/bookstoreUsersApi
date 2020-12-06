package users

import (
	"strings"

	"github.com/igson/bookstoreUsersApi/utils/errors"
)

//User classe de usuário
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate validar campo de cadastro
func (user *User) Validate() *errors.RestErroAPI {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Formato de email inválido.")
	}
	return nil
}
