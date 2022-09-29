package todos

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetTodos(c *gin.Context) {
	todos, status, err := h.Service.GetTodos()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Status Bad Request : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, status, err := h.Service.CreateTodos(req)
	if err != nil {
		log.Println("Error handler Create : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    res,
	})
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	todoId := c.Param("id")
	_, status, err := h.Service.UpdateTodo(todoId)
	if err != nil {
		log.Println("Error handler Update : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "Update data success",
		"data id": todoId,
	})
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	todoId := c.Param("id")
	_, status, err := h.Service.DeleteTodos(todoId)
	if err != nil {
		log.Println("Error handler Delete : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "deleted success",
		"id":      todoId,
	})
}
