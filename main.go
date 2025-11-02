package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	//Создание роутера
	router := gin.Default()

	// API маршруты
	api := router.Group("/api")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "Сервер работает",
			})
		})
	}

	// Статические файлы фронтенда
	staticDir := "./static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		// Если папка static не существует, используем старые шаблоны
		router.LoadHTMLGlob("templates/*")
		router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
		})
	} else {
		// Раздача статических файлов (JS, CSS, изображения и т.д.)
		router.Static("/assets", filepath.Join(staticDir, "assets"))
		router.StaticFile("/favicon.ico", filepath.Join(staticDir, "favicon.ico"))

		// Все остальные маршруты отдаем index.html для SPA роутинга
		router.NoRoute(func(ctx *gin.Context) {
			indexPath := filepath.Join(staticDir, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				ctx.File(indexPath)
			} else {
				ctx.String(http.StatusNotFound, "Frontend not built. Run 'npm run build' in frontend directory")
			}
		})
	}

	//Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080/")
	log.Println("Для разработки фронтенда запустите: cd frontend && npm run dev")
	router.Run(":8080")
}
