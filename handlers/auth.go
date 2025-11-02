package handlers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Форма регистрации
type RegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

// Форма авторизации
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Сервер
type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (server *Server) Register(ctx *gin.Context) {
	var input RegisterInput

	// Бинд формы с моделью при ее получении с фронта
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка обработки формы " + err.Error()})
		return
	}

	// Проверка несовпадения паролей
	if input.Password1 != input.Password2 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "пароли не совпадают"})
		return
	}

	// Запись ввода в модель пользователя
	user := models.User{Username: input.Username, Password: input.Password1}
	if err := user.HashPassword(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка хэширования пароля " + err.Error()})
		return
	}
	user.HtmlEscapeUsername()

	// Запись модели в БД
	if err := server.db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи в БД " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Пользователь создан"})
}

func (server *Server) Login(ctx *gin.Context) {
	// TODO: сделать аутентификацию
}
