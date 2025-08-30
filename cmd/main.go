package main

import (
	"log"
	"ms-user/config"
	"ms-user/db"
	"ms-user/internal/handler"
	"ms-user/internal/repository"
	"ms-user/internal/router"
	"ms-user/internal/service"
	"net/http"
	"os"
)

func main() {
	cfg := config.Load()
	database := db.ConnectDB(cfg)
	logger := log.New(os.Stdout, "[ms-user] ", log.LstdFlags)

	userRepo := repository.NewUserRepositoryImpl(database)
	userService := service.NewUserService(userRepo, logger)
	userHandler := handler.NewUserHandler(userService, logger)
	mux := router.NewRouter(userHandler)

	logger.Println("Starting server on :" + cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		logger.Fatalf("Server error: %v", err)
	}
}
