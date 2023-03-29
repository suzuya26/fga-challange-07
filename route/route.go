package route

import (
	"sesi_8/handler"
	"sesi_8/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	employee := r.Group("/employee")
	{
		employee.POST("", server.CreateEmployee)
	}
	book := r.Group("/book")
	{
		book.POST("", server.CreateBook)
		book.GET("", server.GetAllBook)
		book.GET("/:id", server.GetBookById)
		book.PUT("/:id", server.UpdateBook)
		book.DELETE("/:id", server.DeleteBook)
	}
}
