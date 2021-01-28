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

// GetUser retorna o usuário pelo ID
func GetUser(c *gin.Context) {

	userID, erroID := getUserID(c.Param("user_id"))

	if erroID != nil {
		c.JSON(erroID.StatusCode, erroID)
		return
	}

	user, erroGerUser := services.UserService.GetUser(userID)

	if erroGerUser != nil {
		c.JSON(erroGerUser.StatusCode, erroGerUser)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))

}

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		erroMessage := errors.NewBadRequestError("invalid json error body")
		c.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	createdUser, erro := services.UserService.CreateUser(user)

	if erro != nil {
		c.JSON(erro.StatusCode, erro)
		return
	}

	c.JSON(http.StatusOK, createdUser.Marshall(c.GetHeader("X-Public") == "true"))

}

// UpdateUser realiza a atualização de dados do usuário
func UpdateUser(c *gin.Context) {

	var user users.User

	userID, erroID := getUserID(c.Param("user_id"))

	if erroID != nil {
		c.JSON(erroID.StatusCode, erroID)
		return
	}

	if erro := c.ShouldBindJSON(&user); erro != nil {
		erroMessage := errors.NewBadRequestError("invalid json error body")
		c.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	userUpdate, erroMessage := services.UserService.UpdateUser(isPartial, user)

	if erroMessage != nil {
		c.JSON(erroMessage.StatusCode, erroMessage)
		return
	}

	c.JSON(http.StatusOK, userUpdate.Marshall(c.GetHeader("X-Public") == "true"))

}

// DeleteUser realiza a exclusão do usuário
func DeleteUser(c *gin.Context) {

	userID, erroID := getUserID(c.Param("user_id"))

	if erroID != nil {
		c.JSON(erroID.StatusCode, erroID)
		return
	}

	erro := services.UserService.DeleteUser(userID)

	if erro != nil {
		c.JSON(erro.StatusCode, erro)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

//Search realiza consulta de usuároios
func Search(c *gin.Context) {

	status := c.Query("status")

	users, erro := services.UserService.Search(status)

	if erro != nil {
		c.JSON(erro.StatusCode, erro)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}

func getUserID(userID string) (int64, *errors.RestErroAPI) {
	ID, erro := strconv.ParseInt(userID, 10, 64)
	if erro != nil {
		return 0, errors.NewBadRequestError("ID deve ser número")
	}
	return ID, nil
}
