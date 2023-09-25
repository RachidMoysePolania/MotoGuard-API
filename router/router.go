package router

import (
	"github.com/RachidMoysePolania/MotoGuard-API/controllers"
	"github.com/RachidMoysePolania/MotoGuard-API/helpers"
	"github.com/RachidMoysePolania/MotoGuard-API/middleware"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine = gin.Default()

func Router() *gin.Engine {
	router.GET("/ping", controllers.HealtCheck)

	/*
		GENERAL SCHEMA API
	*/
	api := router.Group("/api")
	v1 := api.Group("/v1")

	/*
		User Group Unprotected
	*/
	user := v1.Group("/users")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)

	/*
		User Routes Group Protected
	*/
	protectedusers := user.Group("", middleware.AuthMiddleware)
	{
		protectedusers.GET("", controllers.GetallUsers)
		protectedusers.GET("/get_user/:id", controllers.GetUserById)
		protectedusers.PUT("/get_user/:id", controllers.UpdateUserById)
		protectedusers.DELETE("/get_user/:id", controllers.DeleteUserById)
	}

	/*
		Protected Log Routes
	*/
	logs := v1.Group("/log", middleware.AuthMiddleware)
	{
		logs.POST("/save", controllers.SaveLogs)
		logs.GET("/:id", controllers.GetLogById)
		logs.GET("", controllers.GetallLogs)
	}
	return router
}

func Server() {
	gin.SetMode(gin.DebugMode)
	srv := Router()
	srv.Use(helpers.Logger())
	srv.Run()
}
