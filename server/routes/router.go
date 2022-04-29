package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/natanista/go-api/controllers"
	"github.com/natanista/go-api/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
main := router.Group("api/v1")
{


	user := main.Group("users")
	{
		user.POST("/", controllers.CreateUser)
	}
	books := main.Group("books", middlewares.Auth())
	{
		books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)


			main.POST("login", controllers.Login)
	}
} 
return router
}