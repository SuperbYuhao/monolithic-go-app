package server

import (
	"monolithic-go-app/internal/controllers"
	"monolithic-go-app/internal/middleware"
	"monolithic-go-app/internal/repositories"
	"monolithic-go-app/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
}

func NewServer(db *gorm.DB) *Server {
	r := gin.Default()

	r.Use(middleware.GinLoggingMiddleware())

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	api := r.Group("/api")
	{
		api.GET("/users", userController.GetAllUsers)
		api.GET("/users/:id", userController.GetUserByID)
		api.POST("/users", userController.CreateUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.DELETE("/users/:id", userController.DeleteUser)
	}

	return &Server{router: r}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
