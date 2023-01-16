package main

import (
	"VS-WS-2022-Backend/controllers"
	"VS-WS-2022-Backend/db"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnvironmentVariables() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

}

func main() {
	loadEnvironmentVariables()
	handleRequests()
}

func handleRequests() {
	DB := db.Init()
	controller := controllers.New(DB)
	router := gin.Default()
	router.GET("/todos", controller.GetAllTodos)
	router.GET("/todos/:todo", controller.GetTodoByString)
	router.POST("/todos", controller.CreateTodo)
	router.POST("/todos/:todo", controller.CreateTodoFromPathvariable)
	router.PUT("/todos", controller.UpdateTodo)
	router.DELETE("/todos/:todo", controller.DeleteTodo)
	router.DELETE("/todos", controller.DeleteAllTodos)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "service is up and running"})
	})

	router.Run()
}
