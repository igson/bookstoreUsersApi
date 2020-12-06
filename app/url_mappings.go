package app

import (
	"github.com/igson/bookstoreUsersApi/controllers/ping"
	"github.com/igson/bookstoreUsersApi/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
