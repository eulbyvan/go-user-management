/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 12:46:30 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package entity

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}