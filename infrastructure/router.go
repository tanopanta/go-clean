package infrastracture

import (
	gin "github.com/gin-gonic/gin"
	"github.com/tanopanta/go-clean/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSQLHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
