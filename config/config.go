/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:34:56 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package config

import "github.com/eulbyvan/go-user-management/pkg/utility"

type Config struct {
	PostgresConnectionString string
	ServerAddress            string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	// Load your configuration values from a configuration file or other sources
	c.PostgresConnectionString = utility.GetEnv("CONNECTION_STRING", "host=localhost port=5432 user=postgres password=postgres dbname=user_management_db sslmode=disable")
	c.ServerAddress = utility.GetEnv("SERVER_ADDRESS", ":8080")
}