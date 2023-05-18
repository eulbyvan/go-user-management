/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 12:47:03 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"github.com/eulbyvan/go-user-management/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(*entity.User) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUser(*entity.User) (*entity.User, error)
	FindUserByID(int64) (*entity.User, error)
	FindUserByUsername(string) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindUserByID(id int64) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAllUser() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}