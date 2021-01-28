package app

import (
	"github.com/gin-gonic/gin"
	"github.com/igson/bookstoreUsersApi/logger"
)

var (
	router = gin.Default()
)

//StartApplication inicializa a aplicação
func StartApplication() {
	mapUrls()
	logger.Info("Inicializando aplicação....")
	router.Run(":8080")
}
