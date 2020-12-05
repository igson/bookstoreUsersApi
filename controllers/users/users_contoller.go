package users

import (
	"fmt"
	"net/http"

	"github.com/igson/bookstoreUsersApi/services"

	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreUsersApi/domain/users"
)

var (
	counter int
)

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return a bad request to caller
		fmt.Println(err.Error())
		return
	}

	createdUser, erroCreated := services.CreateUser(user)

	if erroCreated != nil {
		//TODO: Handle user creation error
		fmt.Println(erroCreated)
		return
	}

	c.JSON(http.StatusCreated, createdUser)

}

// GetUser retorna o usuário pelo ID
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemente-me")
}

// SearchUser realiza a busca de um usuário
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemente-me")
}
