package users

import (
	"fmt"

	"github.com/igson/bookstoreUsersApi/logger"
	"github.com/igson/bookstoreUsersApi/utils/mysqlutils"

	"github.com/igson/bookstoreUsersApi/datasources/mysql/users_db"

	"github.com/igson/bookstoreUsersApi/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?"
	queryUpdateUser = "UPDATE users set first_name=?, last_name=?, email=? WHERE id = ?"
	queryDeleteUser = "DELETE from users WHERE id = ?"
	queryFindUser   = "select id, first_name, last_name, email, date_created, status FROM users WHERE status=?"
)

var (
	usersDB = make(map[int64]*User)
)

//Search retorna lista de usuários
func (user *User) Search(status string) ([]User, *errors.RestErroAPI) {

	stmt, erro := users_db.Database.Prepare(queryFindUser)

	if erro != nil {
		logger.Error("Erro de sintax SQL", erro)
		return nil, errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	rows, erro := stmt.Query(status)

	if erro != nil {
		return nil, errors.NewInternalServerError(erro.Error())
	}

	defer rows.Close()

	usuarios := make([]User, 0)

	for rows.Next() {
		var user User
		if erro := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); erro != nil {
			return nil, mysqlutils.ParserError(erro)
		}
		usuarios = append(usuarios, user)
	}

	if len(usuarios) == 0 {
		return nil, errors.NewNotFoundErro(fmt.Sprintf("Usuário com status %s não encontrado", status))
	}

	return usuarios, nil

}

//GetUser retorna usuário pelo id
func (user *User) GetUser() *errors.RestErroAPI {

	stmt, erro := users_db.Database.Prepare(queryGetUser)

	if erro != nil {
		logger.Error("Erro de sintax SQL", erro)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	resultado := stmt.QueryRow(user.ID)

	if sqlErro := resultado.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); sqlErro != nil {
		logger.Error("Nenhum dado retornado", sqlErro)
		return errors.NewInternalServerError("Nenhum usuário encontrado pro ID fornecido.")
	}

	return nil

}

//Save cadastrar novo usuário
func (user *User) Save() *errors.RestErroAPI {

	stmt, erro := users_db.Database.Prepare(queryInsertUser)

	//query, erro := users_db.Database.Prepare(queryInsertUser)

	if erro != nil {
		logger.Error("Erro de sintax SQL", erro)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	// resultado, erro := query.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	insertResult, saveErro := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErro != nil {
		logger.Error("Erro ao inserir usuário", erro)
		return errors.NewInternalServerError("Database error")
	}

	userID, erro := insertResult.LastInsertId()

	if erro != nil {
		logger.Error("Erro ao retorna um último ID do usuário", erro)
		return errors.NewInternalServerError("Database error")
	}

	user.ID = userID

	return nil

}

//Update cadastrar novo usuário
func (user *User) Update() *errors.RestErroAPI {

	stmt, erro := users_db.Database.Prepare(queryUpdateUser)

	if erro != nil {
		logger.Error("Erro de sintax SQL", erro)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	if _, erro := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID); erro != nil {
		logger.Error("Erro ao atualizar dados do usuário", erro)
		return errors.NewInternalServerError("Database error")
	}

	return nil

}

//Delete deletar usuário
func (user *User) Delete() *errors.RestErroAPI {

	stmt, erro := users_db.Database.Prepare(queryDeleteUser)

	if erro != nil {
		logger.Error("Erro de sintax SQL", erro)
		return errors.NewInternalServerError("Database error")
	}

	defer stmt.Close()

	if _, erro := stmt.Exec(user.ID); erro != nil {
		logger.Error("Erro ao deletar dados do usuário", erro)
		return errors.NewInternalServerError("Database error")
	}

	return nil

}
