package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/goConcurrencyAPI/src/api/controllers/user"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)


func main() {

	router.GET("/users/:idUser", user.GetUser)
	router.Run(port)
}
