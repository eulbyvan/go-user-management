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

type UserRouter struct {
	userHandler UserHandler
	publicRoute *gin.RouterGroup
}

func (u *UserRouter) SetupRouter() {
	users := u.publicRoute.Group("/users")
	{
		//users.Use(Authentication())
		users.POST("/", u.userHandler.InsertUser)
		users.PUT("/:id", u.userHandler.UpdateUser)
		users.DELETE("/:id", u.userHandler.DeleteUser)
		users.GET("/:id", u.userHandler.FindUserByID)
		users.GET("/user", u.userHandler.FindUserByUsername)
		users.GET("/", u.userHandler.FindAllUser)
	}
}

func NewUserRouter(publicRoute *gin.RouterGroup, userUseCase usecase.UserUsecase) {
	userHandler := NewUserHandler(userUseCase)
	rt := UserRouter{
		userHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
