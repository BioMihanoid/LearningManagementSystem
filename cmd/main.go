package main

import (
	"github.com/BioMihanoid/LearningManagementSystem/internal/config"
	"github.com/BioMihanoid/LearningManagementSystem/internal/handlers"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository"
	"github.com/BioMihanoid/LearningManagementSystem/internal/repository/postgres"
	"github.com/BioMihanoid/LearningManagementSystem/internal/service"
	"github.com/BioMihanoid/LearningManagementSystem/pkg"
	"time"
)

func main() {
	// TODO: parse config: godotenv

	// TODO: init logger: logrus

	db, err := postgres.NewPostgresDB(config.Config{
		DbConfig: config.DbConfig{
			PortDb:  "5433",
			Host:    "localhost",
			User:    "lms_user",
			Pass:    "lms_123",
			Dbname:  "lms_db",
			Sslmode: "disable",
		},
	})
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(pkg.Server)
	if err = srv.Start(config.Config{
		ServerConfig: config.ServerConfig{
			PortServ:    "8082",
			Timeout:     4 * time.Second,
			IdleTimeout: 60 * time.Second,
		},
	}, handler.InitRoutes()); err != nil {
	}
}
