package http

import (
	m "github.com/Mubinabd/chat/internal/delivery/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Mubinabd/chat/internal/delivery/docs"
	"github.com/Mubinabd/chat/internal/delivery/http"
)

// @title Chat API Documentation
// @version 1.0
// @description API for Instant Delivery resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *http.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/register", h.RegisterUser).Use(m.Middleware())
	router.POST("/login", h.LoginUser).Use(m.Middleware())
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/reset-password", h.ResetPassword)
	router.GET("/users", h.GetAllUsers).Use(m.JWTMiddleware())

	router.POST("/v1/upload", h.SaveFile)

	user := router.Group("/v1/user").Use(m.JWTMiddleware())
	{
		user.GET("/profiles", h.GetProfile)
		user.PUT("/profiles", h.EditProfile)
		user.PUT("/passwords", h.ChangePassword)
		user.GET("/setting", h.GetSetting)
		user.PUT("/setting", h.EditSetting)
		user.DELETE("/", h.DeleteUser)
	}

	group := router.Group("/v1/groups").Use(m.JWTMiddleware())
	{
		group.POST("/", h.CreateGroup)
		group.GET("/", h.GetAllGroups)
		group.GET("/:id", h.GetGroupByID)
		group.PUT("/:id", h.UpdateGroup)
		group.DELETE("/:id", h.DeleteGroup)
		group.POST("/:groupID/add-member", h.AddMemberToGroup)
		group.DELETE("/messages", h.AddMessageToGroup)
	}

	return router
}
