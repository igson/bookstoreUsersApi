package users

import (
	"net/http"

	"github.com/igson/bookstoreUsersApi/controllers/users"
)

var rotaDeUsuarios = []Rota{
	{
		URI:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             users.GetUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/:users_id",
		Metodo:             http.MethodGet,
		Funcao:             users.SearchUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             users.CreateUser,
		RequerAutenticacao: false,
	},
}
