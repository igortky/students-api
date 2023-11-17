package routes

import (
	"go_gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", controllers.Index)
	r.GET("/students", controllers.IndexStudents)
	r.GET("/students/:id", controllers.ShowStudent)
	r.GET("/students/cpf/:cpf", controllers.ShowStudentByCPF)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.POST("/student", controllers.CreateStudent)
	r.NoRoute(controllers.RouteNotFound)
	r.Run(":8000")
}
