package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func BooksRoute(router *gin.Engine) {

	book := router.Group("/book")
	{
		book.POST("/create", controllers.CreateBook())
		book.GET("/single/:bookID", controllers.GetBook())
		book.GET("/all", controllers.GetBooks())
		book.PUT("/update", controllers.UpdateBook())
		book.GET("/delete/:bookID", controllers.DeleteBook())
	}
}
