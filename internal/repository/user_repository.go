/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 12:47:03 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"

	"github.com/eulbyvan/go-user-management/internal/entity"
)

type UserRepository interface {
	InsertUser(*entity.User) (*entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUser(*entity.User) error
	FindUserByID(int64) (*entity.User, error)
	FindUserByUsername(string) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user *entity.User) (*entity.User, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	stmt, err := r.db.Prepare("UPDATE users SET username = $1, password = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(user *entity.User) error {
	stmt, err := r.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserByID(id int64) (*entity.User, error) {
	var user entity.User
	stmt, err := r.db.Prepare("SELECT id, username, password FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	stmt, err := r.db.Prepare("SELECT id, username, password FROM users WHERE username = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindAllUser() ([]entity.User, error) {
	var users []entity.User
	rows, err := r.db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
