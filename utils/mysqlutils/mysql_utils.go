package mysqlutils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/igson/bookstoreUsersApi/utils/errors"
)

const (
	UniqueEmailConstraint = "email_UNIQUE"
	ErrorNoRows           = "no rows in result set"
)

//ParserError parseamento de erro do Banco de Dados
func ParserError(erro error) *errors.RestErroAPI {

	sqlErro, ok := erro.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(erro.Error(), ErrorNoRows) {
			return errors.NewNotFoundErro("Nenhum registro com o id fornecido")
		}
		return errors.NewInternalServerError("Erro parsing database response")
	}

	switch sqlErro.Number {

	case 1062:
		return errors.NewBadRequestError("Dados inválidos")
	}

	return errors.NewBadRequestError("Erro ao processar requisção")

}
