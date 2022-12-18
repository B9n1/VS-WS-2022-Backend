package controllers

import (
	models "VS-WS-2022-Backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	h.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func (h handler) GetTodoByString(c *gin.Context) {
	var todos []models.Todo
	if err := h.DB.Where("todo = ?", c.Param("todo")).First(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h handler) CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{Todo: input.Todo, Priority: input.Priority}
	h.DB.Create(&todo)

	c.JSON(http.StatusOK, todo)
}

func (h handler) CreateTodoFromPathvariable(c *gin.Context) {

	todo := models.Todo{Todo: c.Param("todo")}
	h.DB.Create(&todo)

	c.JSON(http.StatusOK, todo)
}
func (h handler) UpdateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&input).Updates(input)

	c.JSON(http.StatusOK, input)
}

func (h handler) DeleteTodo(c *gin.Context) {
	var todos []models.Todo
	if err := h.DB.Where("todo = ?", c.Param("todo")).First(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	h.DB.Delete(&todos)

	c.JSON(http.StatusOK, todos)
}

func (h handler) DeleteAllTodos(c *gin.Context) {
	var todos []models.Todo
	h.DB.Where("1 = 1").Delete(&todos)
	h.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}
