/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:22:10 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package server

import (
	"net/http"

	"github.com/eulbyvan/go-user-management/internal/delivery"
	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/eulbyvan/go-user-management/internal/repository"
	"github.com/eulbyvan/go-user-management/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Initialize(connStr string) error {
	// Initialize database connection
	dsn := connStr
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Automatically create the necessary database tables
	db.AutoMigrate(&entity.User{})

	// Initialize repositories and use cases
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Setup router
	s.router = delivery.SetupRouter(userUsecase)

	return nil
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
