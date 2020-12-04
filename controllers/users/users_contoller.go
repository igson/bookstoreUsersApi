package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemente-me")
}

// GetUser retorna o usuário pelo ID
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemente-me")
}

// SearchUser realiza a busca de um usuário
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemente-me")
}
