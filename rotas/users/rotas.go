package rotas

import (
	"github.com/gin-gonic/gin"
	router "github.com/igson/bookstoreUsersApi/app"
)

//Rota mapeamento das rotas da api
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(c *gin.Context)
	RequerAutenticacao bool
}

//Configurar gerar todas as rotas da API
func Configurar(c *gin.Context) *gin.Context {

	rotas := rotaDeUsuarios

	for _, rota := range rotas {
		router.Handle(rota.Metodo, rota.URI, rota.Funcao)
	}

	return router
}
