package users

import (
	"fmt"

	"github.com/igson/bookstoreUsersApi/utils/dateutils"

	"github.com/igson/bookstoreUsersApi/datasources/mysql/users_db"

	"github.com/igson/bookstoreUsersApi/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryGetUser    = "SELECT first_name, last_name, email, date_created FROM users WHERE id = ?"
)

var (
	usersDB = make(map[int64]*User)
)

//GetUser retorna usuário pelo id
func (user *User) GetUser() *errors.RestErroAPI {

	query, erro := users_db.Database.Prepare(queryGetUser)

	if erro != nil {
		return errors.NewInternalServerError(erro.Error())
	}

	defer query.Close()

	resultado, erro := query.QueryRow()

	if erro := resultado.Scan(); erro != nil {

	}

	return nil
}

//Save cadastrar novo usuário
func (user *User) Save() *errors.RestErroAPI {

	query, erro := users_db.Database.Prepare(queryInsertUser)

	if erro != nil {
		return errors.NewInternalServerError(erro.Error())
	}

	defer query.Close()

	user.DateCreated = dateutils.GetNowString()

	resultado, erro := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if erro != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Erro ao cadastrar usuário: %s", erro.Error()))
	}

	userID, erro := resultado.LastInsertId()

	if erro != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Erro ao cadastrar usuário: %s", erro.Error()))
	}

	user.ID = userID

	return nil

}
