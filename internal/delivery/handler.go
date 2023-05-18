/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:04:33 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package delivery

import (
	"strconv"

	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/eulbyvan/go-user-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	InsertUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	FindUserByID(*gin.Context)
	FindUserByUsername(*gin.Context)
	FindAllUser(*gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}

func (u *userHandler) InsertUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.InsertUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (u *userHandler) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.UpdateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (u *userHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.DeleteUser(&entity.User{ID: id})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (u *userHandler) FindUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.FindUserByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (u *userHandler) FindUserByUsername(c *gin.Context) {
	username := c.Param("username")

	result, err := u.userUsecase.FindUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

func (u *userHandler) FindAllUser(c *gin.Context) {
	result, err := u.userUsecase.FindAllUser()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}