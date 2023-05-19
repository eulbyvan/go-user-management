/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:25:56 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package delivery

import (
	"github.com/eulbyvan/go-user-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userUsecase usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	userHandler := NewUserHandler(userUsecase)

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.Register)
			users.PUT("/:id", userHandler.Edit)
			users.DELETE("/:id", userHandler.Unreg)
			users.GET("/:id", userHandler.GetUserByID)
			users.GET("/user", userHandler.GetUserByUsername)
			users.GET("/", userHandler.GetAllUsers)
		}
	}

	return r
}