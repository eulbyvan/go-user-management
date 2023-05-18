/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:02:15 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/eulbyvan/go-user-management/internal/repository"
)

type UserUsecase interface {
	InsertUser(*entity.User) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUser(*entity.User) (*entity.User, error)
	FindUserByID(int64) (*entity.User, error)
	FindUserByUsername(string) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository}
}

func (u *userUsecase) InsertUser(user *entity.User) (*entity.User, error) {
	return u.userRepository.InsertUser(user)
}

func (u *userUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	return u.userRepository.UpdateUser(user)
}

func (u *userUsecase) DeleteUser(user *entity.User) (*entity.User, error) {
	return u.userRepository.DeleteUser(user)
}

func (u *userUsecase) FindUserByID(id int64) (*entity.User, error) {
	return u.userRepository.FindUserByID(id)
}

func (u *userUsecase) FindUserByUsername(username string) (*entity.User, error) {
	return u.userRepository.FindUserByUsername(username)
}

func (u *userUsecase) FindAllUser() ([]entity.User, error) {
	return u.userRepository.FindAllUser()
}