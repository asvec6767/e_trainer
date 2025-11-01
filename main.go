package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Создание роутера
	router := gin.Default()

	// Подключение шаблонов
	router.LoadHTMLGlob("templates/*")

	// Маршруты
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	//Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080/")
	router.Run(":8080")
}
