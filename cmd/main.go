/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Thu May 18 2023 1:47:03 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"log"

	"github.com/eulbyvan/go-user-management/config"
	"github.com/eulbyvan/go-user-management/pkg/server"
)

func main() {
	// Create a new instance of the config
	cfg := config.NewConfig()

	// Load the configuration values
	cfg.Load()

	// Create a new instance of the server
	srv := server.NewServer()

	// Initialize the server with the configuration values
	err := srv.Initialize(cfg.PostgresConnectionString)
	if err != nil {
		log.Fatalf("Failed to initialize the server: %v", err)
	}

	// Start the server
	err = srv.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
