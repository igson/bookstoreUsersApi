package users

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/igson/bookstoreUsersApi/utils/dateutils"

	"github.com/igson/bookstoreUsersApi/datasources/mysql/users_db"

	"github.com/igson/bookstoreUsersApi/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
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

	resultado := query.QueryRow(user.ID)

	if erro := resultado.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); erro != nil {
		fmt.Println(erro)
		return errors.NewInternalServerError(fmt.Sprintf("Erro ao buscar usário com id %d: %s", user.ID, erro.Error()))
	}

	return nil

}

//Save cadastrar novo usuário
func (user *User) Save() *errors.RestErroAPI {

	stmt, erro := users_db.Database.Prepare(queryInsertUser)

	//query, erro := users_db.Database.Prepare(queryInsertUser)

	if erro != nil {
		return errors.NewInternalServerError(erro.Error())
	}

	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	// resultado, erro := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	insertResult, saveErro := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErro != nil {

		sqlErro, ok := saveErro.(*mysql.MySQLError)

		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("Erro ao cadastrar usuário: %s", saveErro.Error()))
		}

		fmt.Println(sqlErro)
		fmt.Println("Number Erro", sqlErro.Number)
		fmt.Println("Message Erro", sqlErro.Message)
		return errors.NewInternalServerError(fmt.Sprintf("Erro ao cadastrar usuário: %s", saveErro.Error()))
	}

	userID, erro := insertResult.LastInsertId()

	if erro != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Erro ao cadastrar usuário: %s", erro.Error()))
	}

	user.ID = userID

	return nil

}
