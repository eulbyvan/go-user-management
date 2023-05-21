/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:04:33 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/eulbyvan/auth-go"
	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/eulbyvan/go-user-management/internal/usecase"
	"github.com/eulbyvan/go-user-management/pkg/utility"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	InsertUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	FindUserByID(*gin.Context)
	FindUserByUsername(*gin.Context)
	FindAllUser(*gin.Context)
	Login(*gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInDb, _ := u.userUsecase.FindUserByUsername(user.Username)
	if userInDb != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	result, err := u.userUsecase.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}

func (u *userHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInDb, err := u.userUsecase.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	result, err := u.userUsecase.UpdateUser(userInDb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *userHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = u.userUsecase.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	err = u.userUsecase.DeleteUser(&entity.User{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("user with id %d has been deleted", id)

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (u *userHandler) FindUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := u.userUsecase.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *userHandler) FindUserByUsername(c *gin.Context) {
	username := c.Query("username")

	result, err := u.userUsecase.FindUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *userHandler) FindAllUser(c *gin.Context) {
	result, err := u.userUsecase.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (u *userHandler) Login(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInDb, err := u.userUsecase.FindUserByUsername(user.Username)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if user.Username != userInDb.Username && user.Password != userInDb.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	secretKey := utility.GetEnv("SECRET_KEY")
	expireTimeInt, err := strconv.Atoi(utility.GetEnv("TOKEN_EXPIRE_TIME_IN_MINUTES"))
	expiresAt := time.Now().Add(time.Minute * time.Duration(expireTimeInt))

	tokenString, err := jwt.GenerateToken(user.ID, expiresAt, []byte(secretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
