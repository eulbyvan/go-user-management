/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:25:10 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key string, v ...any) string {

	// load .env file
	err := godotenv.Load(".env")
   
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	if key != "" {
	  return os.Getenv(key)
	}

	return v[0].(string)
}

func Encrypt(str string) string {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(encryptedPassword)
}