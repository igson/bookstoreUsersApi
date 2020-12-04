package app

import (
	"github.com/gin-gonic/gin"
)

var (
	//Router xxx
	Router = gin.Default()
)

//StartApplication inicializa a aplicação
func StartApplication() {
	//mapUrls()
	GerarRotas()
	Router.Run(":8080")
}
