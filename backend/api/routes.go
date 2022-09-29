package api

import (
	"github.com/gin-contrib/cors"
	"github.com/khalil9022/HelloWorldGoNextJS/todos"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "DELETE", "PUT", "GET"},
	}))

	todosRepo := todos.NewRepository(s.DB)
	todosService := todos.NewService(todosRepo)
	todosHandler := todos.NewHandler(todosService)

	s.Router.GET("/", todosHandler.GetTodos)
	s.Router.POST("/send", todosHandler.CreateTodo)
	s.Router.PUT("/update/:id", todosHandler.UpdateTodo)
	s.Router.DELETE("/delete/:id", todosHandler.DeleteTodo)
}
