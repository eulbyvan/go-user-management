/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:02:15 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/eulbyvan/go-user-management/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(newUser *entity.User) (*entity.User, error)
	Edit(updatedUser *entity.User) (*entity.User, error)
	Unreg(deletedUser *entity.User) error
	GetUserByID(id int64) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	GetAllUsers() ([]entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// Edit implements UserUsecase
func (u *userUsecase) Edit(updatedUser *entity.User) (*entity.User, error) {
	return u.userRepo.UpdateUser(updatedUser)
}

// GetAllUsers implements UserUsecase
func (u *userUsecase) GetAllUsers() ([]entity.User, error) {
	return u.userRepo.FindAllUser()
}

// GetUserByID implements UserUsecase
func (u *userUsecase) GetUserByID(id int64) (*entity.User, error) {
	return u.userRepo.FindUserByID(id)
}

// GetUserByUsername implements UserUsecase
func (u *userUsecase) GetUserByUsername(username string) (*entity.User, error) {
	return u.userRepo.FindUserByUsername(username)
}

// Register implements UserUsecase
func (u *userUsecase) Register(newUser *entity.User) (*entity.User, error) {
	// Business Logic
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser.Password = string(encryptedPassword)

	// Tulis ke db melalui repo
	return u.userRepo.InsertUser(newUser)
}

// Unreg implements UserUsecase
func (u *userUsecase) Unreg(deletedUser *entity.User) error {
	return u.userRepo.DeleteUser(deletedUser)
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}