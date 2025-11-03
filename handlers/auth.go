package handlers

import (
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Форма регистрации
type RegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

// Форма авторизации
type LoginInput struct {
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
	var input LoginInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}

	token, err := server.LoginCheck(user.Username, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка валидации логина и пароля"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (server *Server) LoginCheck(username, password string) (string, error) {
	var err error

	user := models.User{}

	if err = server.db.Model(models.User{}).Where("username=?", username).Take(&user).Error; err != nil {
		return "", err
	}

	err = models.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil

}
