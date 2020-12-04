package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping realizar o retorno com o pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong \n")
}
