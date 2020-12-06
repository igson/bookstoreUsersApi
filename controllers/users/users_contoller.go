package users

import (
	"net/http"
	"strconv"

	"github.com/igson/bookstoreUsersApi/utils/errors"

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
		erroMessage := errors.NewBadRequestError("invalid json error body")
		c.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	createdUser, erro := services.CreateUser(user)

	if erro != nil {
		c.JSON(erro.StatusCode, erro)
		return
	}

	c.JSON(http.StatusCreated, createdUser)

}

// GetUser retorna o usuário pelo ID
func GetUser(c *gin.Context) {

	userID, erro := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if erro != nil {
		erroMessage := errors.NewBadRequestError("ID deve ser número")
		c.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	user, erroGerUser := services.GetUser(userID)

	if erroGerUser != nil {
		c.JSON(erroGerUser.StatusCode, erroGerUser)
		return
	}

	c.JSON(http.StatusOK, user)

}

// SearchUser realiza a busca de um usuário
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemente-me")
}
