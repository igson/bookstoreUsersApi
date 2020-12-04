package app

import (
	"github.com/gin-gonic/gin"
	rotas "github.com/igson/bookstoreUsersApi/rotas/users"
)

//GerarRotas retorna um router com as rotas configuradas
func GerarRotas() *gin.Context {
	r := Router
	return rotas.Configurar(r)
}
